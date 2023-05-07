package payload

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
}
