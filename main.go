package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//cors middleware
	//r.Use(cors.Default())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", " Authorization"}
	r.Use(cors.New(config))

	fmt.Println("hello world")
	r.Run(":3001")
}
