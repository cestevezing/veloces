package router

import (
	"github.com/cestevezing/veloces/internal/controller"
	"github.com/cestevezing/veloces/internal/core/common/middleware"
	"github.com/cestevezing/veloces/internal/core/service"
	"github.com/cestevezing/veloces/internal/infra/repository"
)

func (r *Router) BuildOrderRoutes() {
	orderGroup := r.RouterGroup.Group("/orders")
	orderRepository := repository.NewOrderRepository(r.DB)
	productRepository := repository.NewProductRepository(r.DB)
	orderService := service.NewOrderService(orderRepository, productRepository)
	orderController := controller.NewOrderController(orderService)

	orderGroup.Post("/", middleware.IdempotencyMiddleware(r.RedisClient), orderController.Create)
	orderGroup.Get("/:id", orderController.GetByID)
}
