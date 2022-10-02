package main

import (
	"a2goshell/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Hello world")

	r := gin.Default()
	r.Static("/web", "./web")
	r.GET("/magnet/:link", controller.GetMagnetInfo)

	r.Run(":9999")
}
