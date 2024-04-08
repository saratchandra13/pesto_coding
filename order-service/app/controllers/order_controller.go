package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pesto_coding/order_service/app/services"
	"strconv"
)

type OrderController struct {
	orderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var request struct {
		UserID     uint   `json:"user_id"`
		ProductIDs []uint `json:"product_ids"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order, err := oc.orderService.CreateOrder(request.UserID, request.ProductIDs)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, order)
}

func (oc *OrderController) UpdateOrderStatus(c *gin.Context) {
	orderID, _ := strconv.Atoi(c.Param("order_id"))

	var request struct {
		Status string `json:"status"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order, err := oc.orderService.UpdateOrderStatus(uint(orderID), request.Status)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, order)
}
