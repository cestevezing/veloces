package requests

type OrderCreate struct {
	CustomerName string            `json:"customer_name" validate:"required"`
	Items        []OrderItemCreate `json:"items" validate:"required,min=1,dive"`
}

type OrderItemCreate struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,gt=0"`
}
