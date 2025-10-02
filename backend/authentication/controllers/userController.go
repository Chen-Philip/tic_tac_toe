package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tictactoe/authentication/helpers"
	"tictactoe/authentication/models"
	"tictactoe/database"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string {
	// Password needs to be hashed before being added to the database incase the database is breached, all the passwords
	// arent just available
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // cost is how expensive the hash is, higher = stronger
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email or password is incorrect 2")
		check = false
	}
	return check, msg
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			defer cancel()
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			defer cancel()
			return
		}

		// Counts how many documents with the given field ("username": user.Username)
		count, err := userCollection.CountDocuments(ctx, bson.M{"username": user.Username})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the username"})
			log.Panic(err)
		}

		password := HashPassword(*user.Password)
		user.Password = &password
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this username already exists"})
			return
		}

		// Creates the user
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helpers.GenerateAllTokens(*user.Username, *user.First_name, *user.Last_name, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		// Inserts the user into the database
		_, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, user)
	}

}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			defer cancel()
			return
		}

		// Try to find the user in our database given the username
		err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect 1"})
			return
		}

		// This is just additional checks (not mandatory)
		if foundUser.Username == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		// With the found user, check if the passwords match with each other
		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helpers.GenerateAllTokens(
			*foundUser.Username,
			*foundUser.First_name,
			*foundUser.Last_name,
			*&foundUser.User_id,
		)
		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		err = userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, foundUser)

	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Dont have an admin so dont need this
		//helper.CheckUserTyoe(c, "ADMIN"); err != nil{
		//	c.JSON(http.StatusBadRequest, gin.H{"error", err.Error()})
		//	return
		//}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		// How many records we want per page
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10 // Default value
		}

		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1 // Default value
		}

		// like skip and limimt in node js
		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex")) // startIndex doesn't get changed if URL doesn't have it

		// Matches all documents since it an empty filter
		matchStage := bson.D{{"$match", bson.D{{}}}}

		// Group the documents in the state by a particular id, and then find the counts of the group. THis makes us
		// 	lose the data so we have $push to keep the data
		groupState := bson.D{{"$group", bson.D{
			{"_id", bson.D{{"_id", "null"}}},     // Group all the data based off id
			{"total_count", bson.D{{"$sum", 1}}}, // Create a total count
			// Keep the data in the root without this we will only have a group objext with the count and no data
			{"data", bson.D{{"$push", "$$ROOT"}}},
		}}}

		// Define which data points / fields go to the user / caller and which ones dont
		projectStage := bson.D{{"$project", bson.D{
			{"_id", 0},         // Removes id
			{"total_count", 1}, // Keeps total_count
			{"user_items", bson.D{
				{"$slice", []interface{}{"$data", startIndex, recordPerPage}}, // paginated slice of "data"
			}},
		}}}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupState, projectStage,
		})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		}

		var allUsers []bson.M
		if err = result.All(ctx, &allUsers); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allUsers[0])
	}
}

// Only ADMINS can get this because users shouldn't be able to access other users
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User

		// Mongodb holds the data as jsons, but golang doesnt understand json. Decode helps convert it to the struct
		// 	we made earlier to help golang understnad
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
