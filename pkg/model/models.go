package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       uint32 `json:"price"`
	Garantia    uint32 `json:"garantia"`
	Category    string `json:"category"`
}
type Categories struct {
	Category string `json:"category"`
}

type Customer struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
}
