package service

import (
	"context"
	"../model"
)

type BookingService interface {
	GetAll(ctx context.Context) ([]model.Booking, error)
	Insert(ctx context.Context, body map[string]interface{}) (bool, error)
	GetBookingByUser(ctx context.Context) ([]model.Booking, error)
	GetListHotel(ctx context.Context) ([]model.Hotel, error)
}
