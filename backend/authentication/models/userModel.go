package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User Helps convert json to something that golang can use and vice versa
/*
 *	json: first... tells golang that the name from the database is spelled with a lowercase f, when working with
 * 		datatype is going to be the Firstname with a capital F. validate is required, min 2 chars ...
 *
 * Didn't use here, but validate ..., eq=ADMIN|eq=USER, is the same as an enum
 */
type User struct {
	ID            primitive.ObjectID `bson:bson:"_id`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string            `json:"password" validate:"required,min=6"`
	Username      *string            `json:"username" validate:"required,min=6"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
