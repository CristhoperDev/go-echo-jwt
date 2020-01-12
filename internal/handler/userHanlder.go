package handler

import (
	"encoding/json"
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