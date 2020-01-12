package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go-echo-jwt/internal/dao"
	"go-echo-jwt/internal/model"
	"net/http"
	"time"
)

func RegisterUser(c echo.Context) error {
	var jsonObj model.JsonResult
	var status int

	decoder := json.NewDecoder(c.Request().Body)
	var userBody model.UserBody
	if err := decoder.Decode(&userBody); err != nil {
		jsonObj.Status = http.StatusBadRequest
		jsonObj.Data = "Body not well formed"
		status = http.StatusBadRequest
		return c.JSON(status, jsonObj)
	}

	err := dao.InsertUser(userBody.Username, userBody.Password)


	if err != nil {
		jsonObj.Status = http.StatusBadRequest
		jsonObj.Data = err
		status = http.StatusBadRequest
	} else {
		jsonObj.Status = http.StatusOK
		jsonObj.Data = "User Created"
		status = http.StatusOK
	}

	return c.JSON(status, jsonObj)
}

func generateToken(id int) (map[string]string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	// Create refresh roken
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}

func UserLogin(c echo.Context) error {
	var jsonObj model.JsonResultLogin
	var status int

	decoder := json.NewDecoder(c.Request().Body)
	var userBody model.UserBody
	if err := decoder.Decode(&userBody); err != nil {
		jsonObj.Status = http.StatusBadRequest
		jsonObj.Data = "Body not well formed"
		status = http.StatusBadRequest
		return c.JSON(status, jsonObj)
	}

	data, err := dao.UserLoginByUsernamePassword(userBody.Username, userBody.Password)

	//get token
	getToken,_ := generateToken(data[0].IdUser)

	if len(data) == 0 {
		jsonObj.Status = http.StatusNotFound
		jsonObj.Data = err
		status = http.StatusNotFound
	} else {
		jsonObj.Status = http.StatusOK
		jsonObj.Token = getToken["access_token"]
		jsonObj.RefreshToken = getToken["refresh_token"]
		jsonObj.Data = data
		status = http.StatusOK
	}

	return c.JSON(status, jsonObj)
}

func GetUsers(c echo.Context) error {
	var jsonObj model.JsonResult
	var status int

	data, err := dao.GetAllUsers()

	if len(data) == 0 {
		jsonObj.Status = http.StatusNotFound
		jsonObj.Data = err
		status = http.StatusNotFound
	} else {
		jsonObj.Status = http.StatusOK
		jsonObj.Data = data
		status = http.StatusOK
	}

	return c.JSON(status, jsonObj)
}

func RefreshToken(c echo.Context) error {
	var jsonObj model.JsonResult
	var status int

	decoder := json.NewDecoder(c.Request().Body)
	var refreshTokenBody model.RefreshTokenBody
	if err := decoder.Decode(&refreshTokenBody); err != nil {
		jsonObj.Status = http.StatusBadRequest
		jsonObj.Data = "Body not well formed"
		status = http.StatusBadRequest
		return c.JSON(status, jsonObj)
	}

	token, err := jwt.Parse(refreshTokenBody.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		if int(claims["sub"].(float64)) == 1 {

			newTokenPair, err := generateToken(refreshTokenBody.IdUser)
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, newTokenPair)
		}

		return echo.ErrUnauthorized
	}

	return err
}