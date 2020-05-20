package config

import (
	"../controller"
	"../hanler"
	"../service/impl"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ApplicationContext struct {
	AuthenticationHandler *hanler.AuthenticationHandler
	HotelController       *controller.HotelController
	UserController        *controller.UserController
	BookingController     *controller.BookingController
}

func NewApplicationContext() (*ApplicationContext, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	db := client.Database("Hotel-App")

	authenticationHandler := hanler.NewAuthenticationHandler()

	hotelService := impl.NewHotelServiceImpl(db)
	hotelController := controller.NewHotelController(hotelService)

	userService := impl.NewUserServiceImpl(db)
	userController := controller.NewUserController(userService)

	bookingService := impl.NewBookingServiceImpl(db, hotelService)
	bookingController := controller.NewBookingController(bookingService)

	return &ApplicationContext{
		AuthenticationHandler: authenticationHandler,
		HotelController:       hotelController,
		UserController:        userController,
		BookingController:     bookingController,
	}, err
}
