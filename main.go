package main

import (
	"GinWebServer/routes"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setUpLogger() {
	f, _ := os.Create("ginWebServer.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setUpLogger()
	r := gin.New()
	// , middleware.BasicAuth()
	// r.Use(gin.Recovery(), middleware.CustomLogger()) //middleware registration in gin
	routes.Init(r)
	r.Run(":8000")
}
