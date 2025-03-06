package response

type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Items struct {
	ID       uint    `json:"id"`
	Quantity int     `json:"quantity"`
	Subtotal float64 `json:"subtotal"`
	Product  Product `json:"product"`
}

type Order struct {
	ID           uint    `json:"id"`
	CustomerName string  `json:"customer_name"`
	TotalAmount  float64 `json:"total_amount"`
	Items        []Items `json:"items"`
}
