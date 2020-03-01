package midfunctions

import (
	"gitlab.com/tiktok-server/common"
	"gitlab.com/tiktok-server/databases"
	"gitlab.com/tiktok-server/models"
	"gitlab.com/tiktok-server/utils"
	
)

//User manages functions related to user object
type User struct {
	utils *utils.Utils
}

//Insert is a method unser user object which will validate and sign up a user
func (u *User) Insert(user models.User) (err error) {

	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	err = collection.Insert(&user)
	return

}
