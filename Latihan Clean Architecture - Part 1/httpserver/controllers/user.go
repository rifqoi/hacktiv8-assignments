package controllers

import (
	"fmt"
	"net/http"
	"sesi6-gin/httpserver/controllers/params"
	"sesi6-gin/httpserver/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUser(ctx *gin.Context) {
	var req params.UserCreateRequest

	// step : (2) nge parsing data dari client
	// dan dimasukin ke params
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// step : (3) ngirim data ke services
	response := services.CreateUser(&req)

	// step : (9) kirim ke client
	WriteJsonResponse(ctx, response)
}

func GetUsers(ctx *gin.Context) {
	response := services.GetUsers()
	WriteJsonResponse(ctx, response)
}

func GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id)
	fmt.Println(len(services.Users))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if id > len(services.Users) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("no users with id %d", id),
		})
		return
	}

	fmt.Println("id", id)
	response := services.GetUserByID(id - 1) // id - 1 to get index of slice
	WriteJsonResponse(ctx, response)
}
