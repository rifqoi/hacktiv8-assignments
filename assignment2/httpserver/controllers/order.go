package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/controllers/params"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/controllers/views"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/services"
)

func isValidISODate(date string) error {
	_, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return err
	}
	return nil
}

type OrderController struct {
	svc *services.OrderSVC
}

func NewOrderController(svc *services.OrderSVC) *OrderController {
	return &OrderController{
		svc: svc,
	}
}

func (o *OrderController) CreateOrder(ctx *gin.Context) {
	var req params.OrderCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		WriteJSONResponse(ctx, views.BadRequestError(err, err.Error()))
		return
	}

	err = isValidISODate(req.OrderedAt)
	if err != nil {
		WriteJSONResponse(ctx, views.BadRequestError(err, "Invalid date, please comply to RFC3339 date format!"))
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		WriteJSONResponse(ctx, views.BadRequestError(err, "Items cannot be zero or negative!"))
		return
	}

	response := o.svc.CreateOrder(&req)
	WriteJSONResponse(ctx, response)
}

func (o *OrderController) GetOrders(ctx *gin.Context) {
	limitParams := ctx.Query("limit")
	limit, err := strconv.Atoi(limitParams)
	if err != nil {
		limit = 4
	}
	response := o.svc.GetOrders(limit)

	WriteJSONResponse(ctx, response)
}

func (o *OrderController) UpdateOrder(ctx *gin.Context) {
	var req params.OrderUpdateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		WriteJSONResponse(ctx, views.BadRequestError(err, err.Error()))
		return
	}

	err = isValidISODate(req.OrderedAt)
	if err != nil {
		WriteJSONResponse(ctx, views.BadRequestError(err, "Invalid date, please comply to RFC3339 date format!"))
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		WriteJSONResponse(ctx, views.BadRequestError(err, "Items cannot be zero or negative!"))
		return
	}

	response := o.svc.UpdateOrder(&req)
	WriteJSONResponse(ctx, response)
}

func (o *OrderController) DeleteOrderByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		WriteJSONResponse(ctx, views.BadRequestError(err, "ID must be a number"))
		return
	}

	if id <= 0 {
		WriteJSONResponse(ctx, views.BadRequestError(err, "ID must be greater than 0!"))
	}

	response := o.svc.DeleteOrderByID(id)
	WriteJSONResponse(ctx, response)
}
