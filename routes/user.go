package routes

import (
	"sudut_literat/controllers"
	"sudut_literat/repositories"
	"sudut_literat/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// UserRoutes is a function that defines all the routes related to the user
func UserRoutes(app *gin.RouterGroup, Db *gorm.DB ) {
	// @BasePath /api/v1
	validate :=validator.New()
	repository := repositories.NewUserRepository(Db)
	service := services.NewUserService(repository, validate)
	userCotroller := controllers.NewUserController(service)


	app.GET("/user", controllers.UserIndex)
	app.GET("/user/:id", userCotroller.FindByID)
	app.GET("/user/list", userCotroller.FindAll)
	app.POST("/user/register", userCotroller.Register)
}