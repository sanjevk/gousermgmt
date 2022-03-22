package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sanjevk/mongogo/configs"
	"github.com/sanjevk/mongogo/routes"
)

func main() {
	r := gin.Default()
	configs.ConnectDB()
	routes.UserRoute(r)
	r.Run("localhost:8080")
}
