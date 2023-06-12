package controllers

import (
	"net/http"
	"strconv"
	"sudut_literat/app/request"
	"sudut_literat/app/response"
	"sudut_literat/helper"
	"sudut_literat/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	userService services.UserService
}

// UserIndex godoc
// @Summary user routes
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /user [get]
func UserIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is the user route",
	})
}

func NewUserController(userService services.UserService) *UserController{
	return &UserController{
		userService: userService,
	}
}

// Register godoc
// @Summary register user
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Param example body request.UserRequest true "string"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /user/register [post]
func (controller *UserController) Register(c *gin.Context) {
	userRequest := request.UserRequest{}
	err := c.ShouldBindJSON(&userRequest)
	helper.Error(err)
	
	controller.userService.CreateUser(userRequest)

	res := response.Response{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: userRequest,
	}

	c.JSON(http.StatusOK, res)
}

// FindAll godoc
// @Summary find all user
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]response.UserResponse}
// @Failure 400 {string} string "bad request"
// @Router /user/list [get]
func (controller *UserController) FindAll(c *gin.Context){
	users := controller.userService.GetAll()

	res := response.Response{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: users,
	}

	c.JSON(http.StatusOK, res)
}

// FindByID godoc
// @Summary find user by id
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} response.Response{data=response.UserResponse}
// @Failure 400 {string} string "bad request"
// @Router /user/{id} [get]
func (controller *UserController) FindByID(c *gin.Context){
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	helper.Error(err)
	user := controller.userService.GetDetail(userId)

	res := response.Response{
		Code: http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: user,
	}

	c.JSON(http.StatusOK, res)
}
