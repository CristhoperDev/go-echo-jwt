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

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = data[0].IdUser
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	if len(data) == 0 {
		jsonObj.Status = http.StatusNotFound
		jsonObj.Data = err
		status = http.StatusNotFound
	} else {
		jsonObj.Status = http.StatusOK
		jsonObj.Token = t
		jsonObj.Data = data
		status = http.StatusOK
	}

	return c.JSON(status, jsonObj)
}