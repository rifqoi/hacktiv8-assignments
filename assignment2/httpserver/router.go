package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/controllers"
)

type Router struct {
	router *gin.Engine
	order  *controllers.OrderController
}

func NewRouter(router *gin.Engine, order *controllers.OrderController) *Router {
	return &Router{
		router: router,
		order:  order,
	}
}

func (r *Router) Start(port string) {
	// step :(1) request masuk, request keluar
	r.router.GET("/orders", r.order.GetOrders)
	r.router.DELETE("/orders/:id", r.order.DeleteOrderByID)
	r.router.POST("/orders", r.order.CreateOrder)
	r.router.PUT("/orders", r.order.UpdateOrder)
	r.router.Run(port)
}
