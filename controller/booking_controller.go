package controller

import (
	"../service"
	"context"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type BookingController struct {
	BookingService service.BookingService
}

func NewBookingController(bookingService service.BookingService ) *BookingController{
	return &BookingController{BookingService: bookingService}
}

func (b *BookingController) GetAll(e echo.Context) error {
	authorization:= e.Get("Authorization")
	ctx := context.WithValue(e.Request().Context(), "Authorization", authorization)
	result, err := b.BookingService.GetAll(ctx)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, result)
}

func (b *BookingController) Insert(e echo.Context) error {
	model:= make(map[string]interface{})
	body := json.NewDecoder(e.Request().Body)
	err:= body.Decode(&model)

	startTime, err2 := time.Parse(time.RFC3339, model["startBookingTime"].(string))
	if err2 != nil {
		return e.JSON(http.StatusInternalServerError, err2.Error())
	}
	endTime, err3 := time.Parse(time.RFC3339, model["endBookingTime"].(string))
	if err3 != nil {
		return e.JSON(http.StatusInternalServerError, err3.Error())
	}

	model["startBookingTime"] = startTime
	model["endBookingTime"] = endTime

	if err!= nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	authorization:= e.Get("Authorization")
	ctx := context.WithValue(e.Request().Context(), "Authorization", authorization)

	result, err2 := b.BookingService.Insert(ctx, model)
	if err2 != nil {
		return e.JSON(http.StatusInternalServerError, err2.Error())
	}
	return e.JSON(http.StatusOK, result)
}

func (b *BookingController) GetBookingByUser(e echo.Context) error {
	authorization:= e.Get("Authorization")
	ctx := context.WithValue(e.Request().Context(), "Authorization", authorization)
	result, err := b.BookingService.GetBookingByUser(ctx)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, result)
}

func (b *BookingController) GetListHotel(e echo.Context) error {
	authorization:= e.Get("Authorization")
	ctx := context.WithValue(e.Request().Context(), "Authorization", authorization)
	result, err := b.BookingService.GetListHotel(ctx)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, result)
}

