package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// Not used since I don't have user types
//func CheckUserType(c *gin.Context, role string) (err error) {
//	userType := c.GetString("user_type")
//	err = nil
//	if userType != role {
//		err = errors.New("unauthorized access to this resource")
//	}
//	return err
//}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	//userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	// Users can only access their own user data, since I don't have usertypes, I dont need to check usertype
	// if userType == "USER" && uid != userId {
	if uid != userId {
		err = errors.New("unauthorized access to this resource")
	}

	// err = CheckUserTipe(c, userType)
	return err
}
