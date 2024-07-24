package main

import (
	"calculator/server/api"
	"github.com/gin-gonic/gin"
    _ "calculator/docs"
	"calculator/db"
	// "github.com/swaggo/gin-swagger"
	// swaggerfiles "github.com/swaggo/files"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for Swagger. You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).
// @termsOfService http://swagger.io/terms/
// swagger:meta
func main() {

	db.InitDB()

	router := gin.Default()

	api.Routes(router)

    router.Run(":8000")

}
