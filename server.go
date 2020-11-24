package main

import (
	"fmt"
	"go-stock/api"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.Static("/images", "./uploaded/images")

	api.Setup(route)
	//route.Run()

	//In case of running on Heroku
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("No Port In Heroku")
		route.Run()
	} else {
		fmt.Println("Environment Port : " + port)
		route.Run(":" + port)
	}
}
