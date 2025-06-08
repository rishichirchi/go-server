package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rishichirchi/go-server/database"
	"github.com/rishichirchi/go-server/handler"
)

func main(){
	env_err := godotenv.Load(".env")
	if env_err != nil {
		log.Fatalln("Error loading .env file")
	}
	app := gin.Default()
	database.InitSupabase()
	app.GET("/servers", handler.ListServers)
	app.GET("/servers/:name", handler.GetServer)

	app.Run(":8080")
}