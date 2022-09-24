package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/config"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/controllers"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/repositories/gorm"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/services"
)

func main() {
	db, err := config.ConnectPostgresGORM()
	if err != nil {
		panic(err)
	}
	// Initialize repositories
	orderRepo := gorm.NewOrderRepo(db)

	// Initialize services
	orderSVC := services.NewOrderSVC(orderRepo)

	// Initialize controllers
	orderController := controllers.NewOrderController(orderSVC)

	// Initialize router
	router := gin.Default()
	app := httpserver.NewRouter(router, orderController)
	app.Start(":4000")

}
