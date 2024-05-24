package router

import (
	"daemon/controller"
	"daemon/handler/Model"
	"github.com/gin-gonic/gin"
)

type Router struct {
	controller *controller.Controller
	data       map[string][]Model.Flight
}

func InitRouter(data map[string][]Model.Flight) *Router {
	return &Router{
		data:       data,
		controller: controller.Init(data),
	}
}

func (r *Router) Install(engine *gin.Engine) {
	engine.Use(r.options)
	engine.GET("/", r.healthCheck)
	RB := engine.Group("/api/v1/")
	{
		RB.GET("/flights", r.search)
	}
}
