package router

import (
	"github.com/cestevezing/veloces/internal/controller"
	"github.com/cestevezing/veloces/internal/core/service"
	"github.com/cestevezing/veloces/internal/infra/repository"
)

func (r *Router) BuildProductRoutes() {
	productGroup := r.RouterGroup.Group("/products")
	productRepository := repository.NewProductRepository(r.DB)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	productGroup.Get("/", productController.GetAll)
	productGroup.Get("/:id", productController.GetByID)
	productGroup.Put("/:id/stock", productController.UpdateStock)
}
