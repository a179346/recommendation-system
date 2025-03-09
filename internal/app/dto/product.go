package dto

type Product struct {
	ProductID   int32   `json:"productID"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}
