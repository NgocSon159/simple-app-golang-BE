package controller

import (
	"../model"
	"../service"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController{
	return &UserController{UserService: userService}
}

func (u *UserController) Login(e echo.Context) error {
	var model model.User
	body := json.NewDecoder(e.Request().Body)
	err:= body.Decode(&model)
	if err!= nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err2:= u.UserService.Login(e.Request().Context(), model)
	if err2 != nil {
		return e.JSON(http.StatusInternalServerError, err2.Error())
	}
	return e.JSON(http.StatusOK, result)
}

func (u *UserController) SignUp(e echo.Context) error {
	var model model.User
	body := json.NewDecoder(e.Request().Body)
	err:= body.Decode(&model)
	if err!= nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err2:= u.UserService.SignUp(e.Request().Context(), model)
	if err2 != nil {
		return e.JSON(http.StatusInternalServerError, err2.Error())
	}
	return e.JSON(http.StatusOK, result)
}
