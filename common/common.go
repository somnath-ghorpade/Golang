package common

import (
	"GinWebServer/config"
	"context"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Secrect_key = "29ePuIQIVavQtjkCeYifSxupSsQvSVoJLkHfg2FGkSTaekdztf"
)

func GetMongoConnection() (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// "mongodb://localhost:27017"
	connectionURL := fmt.Sprintf("mongodb://%s:%s", config.IP, config.PORT)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURL))
	// mongodb://user:pass@localhost:27017/?maxPoolSize=100
	if err != nil {
		log.Error("ERROR:Failed to mongo database", err)
		return nil, ctx, cancel, err
	}
	return client, ctx, cancel, nil
}

func DecodeJWTToken(tokenString string) error {

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secrect_key), nil
	})
	if err != nil {
		log.Error("err", err.Error())
		return err
	}

	if !token.Valid {
		log.Error("Invalid Token")
		return err
	}

	return nil
}
