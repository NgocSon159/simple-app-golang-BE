package service

import (
	"../model"
	"context"
)

type UserService interface {
	Login(ctx context.Context, body model.User) (string, error)
	SignUp(ctx context.Context, body model.User) (bool, error)
}


