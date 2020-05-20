package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserId primitive.ObjectID `bson:"_id, omitempty" json:"-"`
	UserName string `bson:"userName" json:"userName"`
	Password string `bson:"password" json:"password"`
	UserType string `bson:"userType" json:"userType"`
}
