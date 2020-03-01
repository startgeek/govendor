package controllers


import (
	"./utils"
	"./midfunctions"
	"./models"
	"./common"

	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	log "github.com/sirupsen/logrus"
)

//User controller wich has functions that can be used againest user object
type User struct {
	utils utils.Utils
	userMidFunctions midfunctions.User
}

//AddUser gonna add a user to db mongo db to be exact
func (u *User) AddUser(ctx *gin.Context)  {
	var addUser models.AddUser

	if err := ctx.ShouldBindJSON(&addUser); err != nil{
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		return
	} 

	if err := addUser.Validate(); err != nil{
		ctx.JSON(http.StatusBadRequest,models.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	user := models.User{bson.NewObjectId(), addUser.Name, addUser.Password}
	err := u.userMidFunctions.Insert(user)

	if err == nil{
		ctx.JSON(http.StatusOK, models.Message{"successfully"})
		log.Debug("Registered a new user" + user.Name)
	}else{
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[Error]", err)
	}
	
}