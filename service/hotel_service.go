package service

import (
	"../model"
	"context"
)

type HotelService interface {
	GetAll(ctx context.Context) ([]model.Hotel, error)
	Insert(ctx context.Context, body map[string]interface{}) (bool, error)
	Update(ctx context.Context, id string, body map[string]interface{}) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
	GetByUser(ctx context.Context) ([]model.Hotel, error)
	GetAllHotel(ctx context.Context) ([]model.Hotel, error)
	GetById(ctx context.Context, id string) (model.Hotel, error)
}
