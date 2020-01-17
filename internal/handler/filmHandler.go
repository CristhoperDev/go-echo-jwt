package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"go-echo-jwt/internal/dao"
	"go-echo-jwt/internal/model"
	"net/http"
)

func FilmPost(c echo.Context) error {
	var jsonObj model.JsonResult
	var status int

	decoder := json.NewDecoder(c.Request().Body)
	var FilmModelPost model.FilmModelPost
	if err := decoder.Decode(&FilmModelPost); err != nil {
		jsonObj.Status = http.StatusBadRequest
		jsonObj.Data = "Body not well formed"
		status = http.StatusBadRequest
		return c.JSON(status, jsonObj)
	}

	err := dao.InsertFilm(FilmModelPost.Title, FilmModelPost.Description)


	if err != nil {
		jsonObj.Status = http.StatusBadRequest
		jsonObj.Data = err
		status = http.StatusBadRequest
	} else {
		jsonObj.Status = http.StatusOK
		jsonObj.Data = "Film Created"
		status = http.StatusOK
	}

	return c.JSON(status, jsonObj)
}
func FilmPut(c echo.Context) error {
	var jsonObj model.JsonResult
	var status int

	decoder := json.NewDecoder(c.Request().Body)
	var FilmModelPut model.FilmModelPut
	if err := decoder.Decode(&FilmModelPut); err != nil {
		jsonObj.Status = http.StatusBadRequest
		jsonObj.Data = "Body not well formed"
		status = http.StatusBadRequest
		return c.JSON(status, jsonObj)
	}

	err := dao.UpdateFilm(FilmModelPut.Title, FilmModelPut.Description, FilmModelPut.IdFilm)


	if err != nil {
		jsonObj.Status = http.StatusBadRequest
		jsonObj.Data = err
		status = http.StatusBadRequest
	} else {
		jsonObj.Status = http.StatusOK
		jsonObj.Data = "Film Updated"
		status = http.StatusOK
	}

	return c.JSON(status, jsonObj)
}