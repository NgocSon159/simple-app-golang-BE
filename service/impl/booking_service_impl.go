package impl

import (
	"../../model"
	"../../service"
	"../../util"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingServiceImpl struct {
	Collection   *mongo.Collection
	HotelService service.HotelService
}

func NewBookingServiceImpl(db *mongo.Database, hotelService service.HotelService) *BookingServiceImpl {
	return &BookingServiceImpl{Collection: db.Collection("Booking"), HotelService: hotelService}
}

func (b *BookingServiceImpl) GetAll(ctx context.Context) ([]model.Booking, error) {
	var arrBooking []model.Booking
	authorization := ctx.Value("Authorization")
	authenticateClaims := authorization.(jwt.MapClaims)
	if authenticateClaims["userType"] == "A" {
		_, err := util.GetAll(ctx, b.Collection, bson.M{}, &arrBooking)
		if err != nil {
			return nil, err
		}
		return arrBooking, nil
	} else {
		return nil, errors.New("You dont have permission!")
	}

}

func (b *BookingServiceImpl) Insert(ctx context.Context, body map[string]interface{}) (bool, error) {
	authorization := ctx.Value("Authorization")
	authenticateClaims := authorization.(jwt.MapClaims)
	body["createBy"] = authenticateClaims["userName"]

	objId, err := primitive.ObjectIDFromHex(body["hotelId"].(string))
	if err != nil {
		return false, err
	}

	body["hotelId"] = objId

	_, err2 := util.Insert(ctx, b.Collection, body)
	if err2 != nil {
		return false, err2
	}
	return true, nil
}

//API for admin-dashboard
func (b *BookingServiceImpl) GetBookingByUser(ctx context.Context) ([]model.Booking, error) {
	arrHotel, err := b.HotelService.GetByUser(ctx)
	if err != nil {
		return nil, err
	}

	arrHotelId := make([]primitive.ObjectID, len(arrHotel))
	for i := 0; i < len(arrHotel); i++ {
		arrHotelId = append(arrHotelId, arrHotel[i].HotelId)
	}

	query := bson.M{"hotelId": bson.M{"$in": arrHotelId}}
	var arrBooking []model.Booking
	_, err2 := util.GetAll(ctx, b.Collection, query, &arrBooking)
	if err2 != nil {
		return nil, err2
	}
	return arrBooking, nil
}

func (b *BookingServiceImpl) GetListHotel(ctx context.Context) ([]model.Hotel, error) {
	arrHotel, err := b.HotelService.GetAllHotel(ctx)
	if err != nil {
		return nil, err
	}
	return arrHotel, nil

}
