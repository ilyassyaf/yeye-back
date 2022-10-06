package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyassyaf/yeyebackend/controllers"
	"github.com/ilyassyaf/yeyebackend/middleware"
	"github.com/ilyassyaf/yeyebackend/services"
)

type CounterRouteController struct {
	counterController controllers.CounterController
}

func NewCounterRouteController(counterController controllers.CounterController) CounterRouteController {
	return CounterRouteController{counterController}
}

func (cr *CounterRouteController) CounterRoute(rg *gin.RouterGroup, uService services.UserService) {
	r := rg.Group("sequence")
	r.Use(middleware.DeserializeUser(uService))

	r.GET("next", cr.counterController.GetNextSequence)
}