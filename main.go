package main

import (
	"net/http"
	"sudut_literat/config"
	"sudut_literat/helper"
	"sudut_literat/initializers"
	"sudut_literat/routes"

	"github.com/gin-gonic/gin"

	docs "sudut_literat/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	initializers.InitEnv()
	docs.SwaggerInfo.BasePath = "/api/v1"
	Db := config.InitDB()



	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// app group route from package routes
	v1 := app.Group("/api/v1");{
		routes.UserRoutes(v1, Db)
	}

	// swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// serve
	err :=http.ListenAndServe("127.0.0.1:8080", app)

	helper.Error(err)
}