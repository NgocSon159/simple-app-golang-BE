package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	HotelId primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber"`
	Avatar string `bson:"avatar" json:"avatar"`
	Photos []string `bson:"photos" json:"photos"`
	CreateBy string `bson:"createBy,omitempty" json:"createBy,omitempty"`
	UpdateBy string `bson:"updateBy,omitempty" json:"updateBy,omitempty"`
}
