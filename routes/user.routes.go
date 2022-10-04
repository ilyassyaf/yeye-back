package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyassyaf/yeyebackend/controllers"
	"github.com/ilyassyaf/yeyebackend/middleware"
	"github.com/ilyassyaf/yeyebackend/services"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {
	r := rg.Group("users")
	r.Use(middleware.DeserializeUser(userService))

	r.GET("/me", uc.userController.GetMe)
}