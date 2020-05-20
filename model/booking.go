package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Booking struct {
	BookingId primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	HotelId primitive.ObjectID `bson:"hotelId,omitempty" json:"hotelId"`
	Name string `bson:"name,omitempty" json:"name"`
	Description string `bson:"description,omitempty" json:"description"`
	StartBookingTime time.Time `bson:"startBookingTime,omitempty" json:"startBookingTime"`
	EndBookingTime time.Time `bson:"endBookingTime,omitempty" json:"endBookingTime"`
	CreateBy string `bson:"createBy,omitempty" json:"createBy"`
}
