package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error){
	userType := c.GetString("user_type")
	err = nil 
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userid string) (err error){
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil
	if userType != "USER" && uid != userid {
		err = errors.New("Unauthorized to acess this resourse.")
		return err
	}
	return err
}