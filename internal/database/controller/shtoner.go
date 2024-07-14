package controller

import (
	"context"

	db "github.com/MukundSinghRajput/url-shortner/internal/database"
	"github.com/MukundSinghRajput/url-shortner/internal/database/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var colshort = db.Database.Collection("short")

func Add(short model.Shotner) bool {
	if _, found := getByOrigin(short.Origin); found {
		return true
	}

	if _, ok := Get(short.Short); ok {
		return true
	}

	go func() {
		_, err := colshort.InsertOne(context.Background(), short)
		if err != nil {
			return
		}
	}()

	return true
}

func Get(short string) (string, bool) {
	var result model.Shotner
	err := colshort.FindOne(context.Background(), bson.M{"short": short}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", false
		}
		return "", false
	}

	return result.Origin, true
}

func getByOrigin(origin string) (model.Shotner, bool) {
	var result model.Shotner
	err := colshort.FindOne(context.Background(), bson.M{"origin": origin}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Shotner{}, false
		}
		return model.Shotner{}, false
	}
	return result, true
}
