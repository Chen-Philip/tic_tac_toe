package helpers

import (
	"chess/database"
	"context"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/openpgp/packet"

	"log"
	"os"
	"time"
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
