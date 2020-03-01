package main

import (
	"io"
	"os"
	"gitlab.com/startengine/tiktok-server/src/user-microservice/common"
	"gitlab.com/startengine/tiktok-server/src/user-microservice/databases"
	"gitlab.com/startengine/tiktok-server/src/user-microservice/controllers"

	"github.com/gin-gonic/gin"
)

// Main manages main golang application
type Main struct {
	router *gin.Engine
}


func (m *Main) initServer() error {
	var err error
	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize User database
	err = databases.Database.Init()
	if err != nil {
		return err
	}

	// Setting Gin Logger
	if common.Config.EnableGinFileLog {
		f, _ := os.Create("logs/gin.log")
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	m.router = gin.Default()

	return nil
}

func main()  {
	m := Main{}

	// Initialize server
	if m.initServer() != nil {
		println("im in this")
		return
	}

	defer databases.Database.Close()

	c := controllers.User{}
	v1 := m.router.Group("/api/v1")
	{

		user := v1.Group("/users")
		{
			user.POST("", c.AddUser)

		}

	}
	m.router.Run(common.Config.Port)
}