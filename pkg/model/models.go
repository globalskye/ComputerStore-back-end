package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Product struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Image       string `json:"image,omitempty"`
	Description string `json:"description,omitempty"`
	Price       uint32 `json:"price,omitempty"`
	Garantie    uint32 `json:"garantia,omitempty"`
	Category    string `json:"_category,omitempty"`
}
type ProductCategory struct {
	Category string `json:"category"`
}

type Customer struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
}
