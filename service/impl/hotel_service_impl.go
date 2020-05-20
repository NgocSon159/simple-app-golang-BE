package impl

import (
	"../../model"
	"../../util"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelServiceImpl struct {
	Collection *mongo.Collection
}

func NewHotelServiceImpl(db *mongo.Database) *HotelServiceImpl {
	return &HotelServiceImpl{Collection: db.Collection("Hotel")}
}

func (h *HotelServiceImpl) GetAll(ctx context.Context) ([]model.Hotel, error) {
	//var arrHotel []model.Hotel
	//_, err:= util.GetAll(ctx, h.Collection ,bson.M{}, &arrHotel)
	//if err != nil {
	//	return arrHotel, err
	//}
	//return arrHotel, nil
	authorization:= ctx.Value("Authorization")
	authenticateClaims := authorization.(jwt.MapClaims)

	var arrHotel []model.Hotel

	if authenticateClaims["userType"] == "A" {
		_, err:= util.GetAll(ctx, h.Collection, bson.M{}, &arrHotel)
		if err != nil {
			return nil, err
		}
		return arrHotel, nil
	} else {
		query:= bson.M{"createBy": authenticateClaims["userName"]}
		_, err:= util.GetAll(ctx, h.Collection, query, &arrHotel)
		if err != nil {
			return nil, err
		}
		return arrHotel, nil
	}

}

func (h *HotelServiceImpl) Insert(ctx context.Context, body map[string]interface{}) (bool, error) {
	authorization:= ctx.Value("Authorization")
	authenticateClaims := authorization.(jwt.MapClaims)
	body["_id"] = primitive.NewObjectID()
	body["createBy"] = authenticateClaims["userName"]
	_, err:= util.Insert(ctx, h.Collection, body)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (h *HotelServiceImpl) Update(ctx context.Context, id string, body map[string]interface{}) (bool, error) {
	authorization:= ctx.Value("Authorization")
	authenticateClaims := authorization.(jwt.MapClaims)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	body["updateBy"] = authenticateClaims["userName"]
	query := bson.M{"_id" :objId}

	var hotelBefore model.Hotel
	_, err2 := util.GetById(ctx, h.Collection, query, &hotelBefore)
	if err2 != nil {
		return false, err2
	}

	flag:= false
	if hotelBefore.CreateBy == authenticateClaims["userName"] {
		flag = true
	} else if hotelBefore.UpdateBy == authenticateClaims["userName"] {
		flag = true
	}

	if flag == true {
		_, err3:= util.Update(ctx, h.Collection, query, body)
		if err3 != nil {
			return false, err3
		}
		return true, nil
	} else {
		return false, errors.New("You dont have permission!")
	}

}

func (h *HotelServiceImpl) Delete(ctx context.Context, id string) (bool, error) {
	authorization:= ctx.Value("Authorization")
	authenticateClaims := authorization.(jwt.MapClaims)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	query := bson.M{"_id" :objId}

	var hotelBefore model.Hotel
	_, err2 := util.GetById(ctx, h.Collection, query, &hotelBefore)
	if err2 != nil {
		return false, err2
	}

	flag:= false
	if hotelBefore.CreateBy == authenticateClaims["userName"] {
		flag = true
	} else if hotelBefore.UpdateBy == authenticateClaims["userName"] {
		flag = true
	}

	if flag == true {
		_, err2 := util.Delete(ctx, h.Collection, query)
		if err2 != nil {
			return false, err2
		}
	}
	return true, nil
}

func (h *HotelServiceImpl) GetByUser(ctx context.Context) ([]model.Hotel, error) {
	authorization:= ctx.Value("Authorization")
	authenticateClaims := authorization.(jwt.MapClaims)

	var arrHotel []model.Hotel
	query:= bson.M{"createBy": authenticateClaims["userName"]}
	_, err:= util.GetAll(ctx, h.Collection, query, &arrHotel)
	if err != nil {
		return nil, err
	}
	return arrHotel, nil
}

func (h *HotelServiceImpl) GetAllHotel(ctx context.Context) ([]model.Hotel, error) {
	var arrHotel []model.Hotel
	_, err:= util.GetAll(ctx, h.Collection ,bson.M{}, &arrHotel)
	if err != nil {
		return arrHotel, err
	}
	return arrHotel, nil
}

func (h *HotelServiceImpl) GetById(ctx context.Context, id string) (model.Hotel, error) {
	var hotel model.Hotel
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return hotel, err
	}
	_, err2:= util.GetById(ctx, h.Collection ,bson.M{"_id": objId}, &hotel)
	if err2 != nil {
		return hotel, err2
	}
	return hotel, nil
}
