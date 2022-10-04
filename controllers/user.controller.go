package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyassyaf/yeyebackend/models"
	"github.com/ilyassyaf/yeyebackend/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService}
}

func (uc *UserController) GetMe(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(*models.DBResponse)

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": models.FilteredResponse(currentUser)}})
}