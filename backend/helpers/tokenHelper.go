package helpers

import (
	"chess/database"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// THe JWT (JSON web token) token makes a token by hashing the data you give it
// it is a compact, url safe way to send information between parties as a json object
// consists of 3 partsL Header, payload signature
// 1. header contains info about the token itself (type, hashing function)
// 2. payload: the actual data of the token (in this case the signed details)
// 3. signature: hash of header + payload + secret key to make sure token wasnt tampered with
type SignedDetails struct {
	Username   string
	First_name string
	Last_name  string
	Uid        string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY") // can add to out .env file

func GenerateAllTokens(username, first_name, last_name, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Username:   username,
		First_name: first_name,
		Last_name:  last_name,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			// Every token needs an expire time.
			// 	Unix converts the time into a unix timestamp (int of seconds since ...)
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			// Expires a week from now
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D // ordered slice of key-value pairs,
	// M is unordered for mondodb, d is ordered

	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", updatedAt})

	upsert := true // if doc exists -> update it, else create a new doc with the given data
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	// We have a _ since we dont care what it returns, we only care that it updates
	_, err := userCollection.UpdateOne(
		ctx,
		filter,                      // finds document with user_id == userId
		bson.D{{"$set", updateObj}}, //
		&opt,                        // makes update to allow for upserts (creating new docs)
	)

	defer cancel()
	if err != nil {
		log.Panic(err)
	}

	return
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	// Claims are the information (KV pairs) stored inside the token's payload
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{}, // token gets decoded into this struct
		func(token *jwt.Token) (interface{}, error) { // tells the parser what secret key to use
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	// Extract the claims and checks if its the right type
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}
