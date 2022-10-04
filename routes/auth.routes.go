package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyassyaf/yeyebackend/controllers"
	"github.com/ilyassyaf/yeyebackend/middleware"
	"github.com/ilyassyaf/yeyebackend/services"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup, userService services.UserService) {
	r := rg.Group("/auth")

	r.POST("/register", rc.authController.SignUpUser)
	r.POST("/login", rc.authController.SignInUser)
	r.GET("/refresh", rc.authController.RefreshAccessToken)
	r.GET("/logout", middleware.DeserializeUser(userService), rc.authController.LogoutUser)
}

