package controller

import (
	"strconv"

	"github.com/cestevezing/veloces/internal/core/dto/requests"
	"github.com/cestevezing/veloces/internal/core/dto/response"
	"github.com/cestevezing/veloces/internal/core/port/service"
	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	service service.IOrder
}

func NewOrderController(service service.IOrder) *OrderController {
	return &OrderController{service: service}
}

// @Summary Get order by ID
// @Description Retrieves an order using its unique ID.
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} response.Response "Order retrieved successfully"
// @Failure 400 {object} response.Response "Invalid order ID"
// @Router /api/orders/{id} [get]
func (ctx *OrderController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid order ID")
	}
	order, err := ctx.service.GetByID(idInt)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, "", order)
}

// @Summary Create a new order
// @Description Creates a new order with the provided details.
// @Tags Orders
// @Accept json
// @Produce json
// @Param Idempotency-Key header string true "Idempotency Key to prevent duplicate requests"
// @Param body body requests.OrderCreate true "Order creation payload"
// @Success 201 {object} model.Order "Order created successfully"
// @Failure 400 {object} response.Response "Invalid request body"
// @Router /api/orders [post]
func (ctx *OrderController) Create(c *fiber.Ctx) error {
	var order *requests.OrderCreate
	if err := c.BodyParser(&order); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	newOrder, err := ctx.service.Create(order)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	c.Locals("response", newOrder)
	return response.Success(c, "Order created successfully", newOrder)
}
