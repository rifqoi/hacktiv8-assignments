package services

import (
	"strings"
	"time"

	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/controllers/params"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/controllers/views"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/repositories"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/repositories/models"
	"gorm.io/gorm"
)

type OrderSVC struct {
	orderRepo repositories.OrderRepo
}

func NewOrderSVC(orderRepo repositories.OrderRepo) *OrderSVC {
	return &OrderSVC{
		orderRepo: orderRepo,
	}
}

func (o *OrderSVC) CreateOrder(req *params.OrderCreateRequest) *views.Response {
	// Parse OrderedAt to datetime format
	t, _ := time.Parse(time.RFC3339, req.OrderedAt)

	// Putting it together in order models struct
	items := parseItemsFromCreateReq(req)
	order := &models.Order{
		CustomerName: req.CustomerName,
		OrderedAt:    t,
		Items:        *items,
	}

	// Insert data to database
	err := o.orderRepo.CreateOrder(order)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return views.DataConflict(err)
		}
		return views.InternalServerError(err)
	}
	return views.SuccessCreateResponse(&req, "success")
}

func (o *OrderSVC) GetOrders(limit int) *views.Response {
	orders, err := o.orderRepo.GetOrders(limit)
	if err != nil {
		return views.InternalServerError(err)
	}
	return views.SuccessResponse(parseGetOrders(orders), "SUCCESS_GET_ORDERS")
}

func (o *OrderSVC) UpdateOrder(req *params.OrderUpdateRequest) *views.Response {
	// Parse OrderedAt to datetime format

	items := parseItemsFromUpdateReq(req)

	t, _ := time.Parse(time.RFC3339, req.OrderedAt)
	order := models.Order{
		ID:           req.OrderID,
		CustomerName: req.CustomerName,
		OrderedAt:    t,
		Items:        *items,
	}
	err := o.orderRepo.UpdateOrder(&order)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.FailedNotFound(err)
		}
		return views.InternalServerError(err)
	}

	return views.SuccessResponse(order, "SUCCESS_UPDATE_ORDER")
}

func (o *OrderSVC) DeleteOrderByID(id int) *views.Response {
	var order models.Order
	err := o.orderRepo.DeleteOrderByID(id, &order)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.FailedNotFound(err)
		}
		return views.InternalServerError(err)
	}

	v := &views.OrderDelete{
		ID:           order.ID,
		CustomerName: order.CustomerName,
	}

	return views.SuccessResponse(v, "SUCCESS_DELETE_ORDER")
}

func parseItemsFromUpdateReq(req *params.OrderUpdateRequest) *[]models.Item {
	var items []models.Item
	for _, item := range req.Items {
		a := models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		}
		items = append(items, a)
	}
	return &items
}

func parseItemsFromCreateReq(req *params.OrderCreateRequest) *[]models.Item {
	var items []models.Item
	for _, item := range req.Items {
		a := models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		}
		items = append(items, a)
	}
	return &items
}

func parseGetOrders(orders *[]models.Order) *[]views.OrderGet {
	var orderViews []views.OrderGet
	for _, order := range *orders {
		items := []views.ItemsOrderGet{}
		for _, item := range order.Items {
			items = append(items, views.ItemsOrderGet{
				ID:          item.ID,
				ItemCode:    item.ItemCode,
				Description: item.Description,
				Quantity:    item.Quantity,
			})
		}

		orderViews = append(orderViews, views.OrderGet{
			ID:           order.ID,
			CustomerName: order.CustomerName,
			OrderedAt:    order.OrderedAt,
			Items:        items,
		})
	}

	return &orderViews
}
