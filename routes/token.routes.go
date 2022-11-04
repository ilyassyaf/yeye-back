package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyassyaf/yeyebackend/controllers"
	// "github.com/ilyassyaf/yeyebackend/middleware"
	// "github.com/ilyassyaf/yeyebackend/services"
)

type TokenRouteController struct {
	tokenController controllers.TokenCotroller
}

func NewTokenRouteController(tokenController controllers.TokenCotroller) TokenRouteController {
	return TokenRouteController{tokenController}
}

func (tc *TokenRouteController) TokenRoute(rg *gin.RouterGroup /* , uService services.UserService */) {
	r := rg.Group("token")

	r.GET("/metadata/:id", tc.tokenController.GetMetadata)

	// r.Use(middleware.DeserializeUser(uService))

	r.GET("/all", tc.tokenController.GetAll)
	r.GET("/get", tc.tokenController.Get)
	r.POST("/store", tc.tokenController.Store)
	r.POST("/image-store", tc.tokenController.StoreTokenImage)
	r.POST("/category-store", tc.tokenController.StoreCategory)
	r.GET("/category-all", tc.tokenController.GetAllByCategory)
	r.GET("/category-get", tc.tokenController.GetByCategory)
}
