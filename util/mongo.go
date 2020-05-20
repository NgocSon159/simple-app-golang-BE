package util

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAll(ctx context.Context, collection *mongo.Collection, query interface{}, result interface{}) (bool ,error) {
	cur, err:= collection.Find(ctx, query)
	if err != nil {
		return false, err
	}
	err2:= cur.All(ctx, result)
	if err2 != nil {
		return false, err2
	}
	return true, nil
}

func Insert(ctx context.Context, collection *mongo.Collection, model interface{}) (bool ,error) {
	_, err:= collection.InsertOne(ctx, model)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Update(ctx context.Context, collection *mongo.Collection, query interface{}, model interface{}) (bool ,error) {
	update:= bson.M{
		"$set" : model,
	}
	_, err:= collection.UpdateOne(ctx, query, update)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Delete(ctx context.Context, collection *mongo.Collection, query interface{}) (bool ,error) {
	_, err:= collection.DeleteOne(ctx, query)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetById(ctx context.Context, collection *mongo.Collection, query interface{}, result interface{}) (bool, error) {
	results := collection.FindOne(ctx, query)
	err := results.Decode(result)
	if err != nil {
		return false, err
	}
	return true, nil
}




