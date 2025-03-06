package controller

import (
	"net/http"
	"strconv"

	"github.com/cestevezing/veloces/internal/core/common/utils"
	"github.com/cestevezing/veloces/internal/core/dto/requests"
	"github.com/cestevezing/veloces/internal/core/dto/response"
	"github.com/cestevezing/veloces/internal/core/port/service"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	service service.IProduct
}

func NewProductController(service service.IProduct) *ProductController {
	return &ProductController{service: service}
}

// @Summary Get All Products
// @Description Retrieve a list of all products.
// @Tags Products
// @Produce json
// @Success 200 {object} []model.Product "Products retrieved successfully"
// @Failure 404 {object} response.Response "Products not found"
// @Router /api/products [get]
func (ctr *ProductController) GetAll(c *fiber.Ctx) error {
	products, err := ctr.service.GetAll(c.Context())
	if err != nil {
		return response.Error(c, http.StatusNotFound, err.Error())
	}
	return response.Success(c, "Products retrieved successfully", products)
}

// @Summary Get Product by ID
// @Description Retrieve a single product by its ID.
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Product "Product retrieved successfully"
// @Failure 400 {object} response.Response "Invalid product ID"
// @Router /api/products/{id} [get]
func (ctr *ProductController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid product ID")
	}
	product, err := ctr.service.GetByID(c.Context(), idInt)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}
	return response.Success(c, "Product retrieved successfully", product)
}

// @Summary Update Product Stock
// @Description Update the stock quantity of a product.
// @Tags Products
// @Accept json
// @Produce json
// @Param Idempotency-Key header string true "Idempotency Key to prevent duplicate requests"
// @Param id path int true "Product ID"
// @Param body body requests.ProductStock true "Stock update payload"
// @Success 200 {object} model.Product "Product stock updated successfully"
// @Failure 400 {object} response.Response "Invalid request body or product ID"
// @Router /api/products/{id}/stock [put]
func (ctr *ProductController) UpdateStock(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid product ID")
	}
	var productStock *requests.ProductStock
	err = c.BodyParser(&productStock)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body")
	}
	if err := utils.ValidateStruct(productStock); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}
	result, err := ctr.service.UpdateStock(c.Context(), idInt, productStock)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}
	c.Locals("response", result)
	return response.Success(c, "Product stock updated successfully", result)
}
