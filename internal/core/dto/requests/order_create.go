package requests

type OrderCreate struct {
	CustomerName string            `json:"customer_name"`
	Items        []OrderItemCreate `json:"items"`
}

type OrderItemCreate struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
