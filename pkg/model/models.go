package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Adress   string `json:"adress"`
	Role     string `json:"role" `
}
type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       uint32 `json:"price"`
	Garantia    uint32 `json:"garantia"`
	Category    string `json:"category"`
	Provider    string `json:"provider"`
}
type UserCard struct {
	ID     int `json:"ID"`
	ItemId int `json:"ItemId"`
}
type Categories struct {
	Category string `json:"category"`
}
type Providers struct {
	Provider string `json:"provider"`
}

type Customer struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
}
