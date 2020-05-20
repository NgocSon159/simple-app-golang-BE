package impl

import (
	"../../model"
	"../../util"
	"context"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserServiceImpl struct {
	Collection *mongo.Collection
}

func NewUserServiceImpl(db *mongo.Database) *UserServiceImpl {
	return &UserServiceImpl{Collection: db.Collection("User")}
}

func (u *UserServiceImpl) Login(ctx context.Context, body model.User) (string, error) {
	query := make(map[string]string)
	query["userName"] = body.UserName
	var user model.User
	_, err := util.GetById(ctx, u.Collection, query, &user)
	if err != nil {
		return "", err
	}
	if user.UserName != "" {
		err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if err2 != nil {
			return "false", err2
		}
		token, err3 := createToken(user)
		if err3 != nil {
			return "", err3
		}
		return token, nil
	}
	return "", nil
}

func createToken(user model.User) (string, error) {
	claims:= jwt.MapClaims{}
	claims["userName"] = user.UserName
	claims["userType"] = user.UserType
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserServiceImpl) SignUp(ctx context.Context, body model.User) (bool, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		return false, err
	}
	body.Password = string(bytes)
	body.UserId = primitive.NewObjectID()
	body.UserType = "C"
	_, err2 := util.Insert(ctx, u.Collection, body)
	if err2 != nil {
		return false, err2
	}
	return true, nil
}
