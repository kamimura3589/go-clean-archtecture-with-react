package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Println("hello world")
	r.Run(":3001")
}
