package route

import (
	"../config"
	"github.com/labstack/echo"
)

type HotelAppRoute struct {
	Echo *echo.Echo
}

func NewHotelAppRoute(e *echo.Echo) (*HotelAppRoute, error) {
	applicationContext, _:= config.NewApplicationContext()
	a:= applicationContext.AuthenticationHandler

	hotelController:= applicationContext.HotelController
	hotelPath:= "/hotel"
	hotelGroup:= e.Group(hotelPath, a.Authenticate)
	hotelGroup.GET("", hotelController.GetAll)
	hotelGroup.GET("/:id", hotelController.GetById)
	hotelGroup.POST("", hotelController.Insert)
	hotelGroup.PUT("/:id", hotelController.Update)
	hotelGroup.DELETE("/:id", hotelController.Delete)

	userController:= applicationContext.UserController
	userPath:= "/user"
	e.POST(userPath + "/login", userController.Login)
	e.POST(userPath + "/signUp", userController.SignUp)

	bookingController:= applicationContext.BookingController
	bookingPath:= "/booking"
	bookingGroup:= e.Group(bookingPath, a.Authenticate)
	bookingGroup.GET("", bookingController.GetAll)
	bookingGroup.GET("/hotel", bookingController.GetListHotel)
	bookingGroup.POST("", bookingController.Insert)
	bookingGroup.GET("/user", bookingController.GetBookingByUser)
	return &HotelAppRoute{Echo: e}, nil

}
