package requests

type ProductStock struct {
	NewStock int `json:"new_stock" validate:"required,gt=0"`
}
