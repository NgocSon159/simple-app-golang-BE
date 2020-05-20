package controller

import (
	"../service"
	"context"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

type HotelController struct {
	HotelService service.HotelService
}

func NewHotelController(hotelService service.HotelService) *HotelController{
	return &HotelController{HotelService: hotelService}
}

func (h *HotelController) GetAll(e echo.Context) error {
	authorization:= e.Get("Authorization")
	ctx := context.WithValue(e.Request().Context(), "Authorization", authorization)
	result, err := h.HotelService.GetAll(ctx)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, result)
}

func (h *HotelController) Insert(e echo.Context) error {
	model:= make(map[string]interface{})
	body := json.NewDecoder(e.Request().Body)
	err:= body.Decode(&model)
	if err!= nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	authorization:= e.Get("Authorization")
	ctx := context.WithValue(e.Request().Context(), "Authorization", authorization)

	result, err2 := h.HotelService.Insert(ctx, model)
	if err2 != nil {
		return e.JSON(http.StatusInternalServerError, err2.Error())
	}
	return e.JSON(http.StatusOK, result)
}

func (h *HotelController) Update(e echo.Context) error {
	model:= make(map[string]interface{})
	body := json.NewDecoder(e.Request().Body)
	err:= body.Decode(&model)
	if err!= nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	id:= e.Param("id")

	authorization:= e.Get("Authorization")
	ctx := context.WithValue(e.Request().Context(), "Authorization", authorization)

	result, err2 := h.HotelService.Update(ctx, id, model)
	if err2 != nil {
		return e.JSON(http.StatusInternalServerError, err2.Error())
	}
	return e.JSON(http.StatusOK, result)
}

func (h *HotelController) Delete(e echo.Context) error {
	id:= e.Param("id")
	authorization:= e.Get("Authorization")
	ctx := context.WithValue(e.Request().Context(), "Authorization", authorization)
	result, err2 := h.HotelService.Delete(ctx, id)
	if err2 != nil {
		return e.JSON(http.StatusInternalServerError, err2.Error())
	}
	return e.JSON(http.StatusOK, result)
}

func (h *HotelController) GetById(e echo.Context) error {
	id:= e.Param("id")
	authorization:= e.Get("Authorization")
	ctx := context.WithValue(e.Request().Context(), "Authorization", authorization)
	result, err2 := h.HotelService.GetById(ctx, id)
	if err2 != nil {
		return e.JSON(http.StatusInternalServerError, err2.Error())
	}
	return e.JSON(http.StatusOK, result)
}


