package model

import "time"

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
	CategoryId  uint32 `json:"categoryId"`
	Category    string `json:"category"`
	ProviderId  uint32 `json:"providerId"`
	Provider    string `json:"provider"`
}
type ProductToAdd struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Count       int    `json:"count"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       uint32 `json:"price"`
	Garantia    uint32 `json:"garantia"`
	CategoryId  uint32 `json:"categoryId"`
	Category    string `json:"category"`
	ProviderId  uint32 `json:"providerId"`
	Provider    string `json:"provider"`
}

type Stock struct {
	Count       int    `json:"count"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       uint32 `json:"price"`
	Garantia    uint32 `json:"garantia"`
	Category    string `json:"category"`
	Provider    string `json:"provider"`
}
type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Login     string `json:"login"`
	Passwd    string `json:"passwd"`
	Phone     string `json:"phone"`
	WorkTime  string `json:"worktime"`
	Salary    string `json:"salary"`
}
type _ struct {
	ID     int                 `json:"id"`
	UserId int                 `json:"userId"`
	Items  []UserCardItemCount `json:"items"`
}
type UserCard struct {
	Id         int    `json:"id"`
	UserId     int    `json:"userId"`
	UserAdress string `json:"userAdress"`
	Items      []struct {
		Quantity    int    `json:"quantity"`
		TotalPrice  int    `json:"totalPrice"`
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Image       string `json:"image"`
		Description string `json:"description"`
		Price       int    `json:"price"`
		Garantia    int    `json:"garantia"`
		Category    string `json:"category"`
		Provider    string `json:"provider"`
	} `json:"items"`
}

type UserCardItemCount struct {
	Count      uint32  `json:"count"`
	TotalPrice uint32  `json:"totalPrice"`
	Item       Product `jso	n:"item"`
}
type Categories struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
}
type DashBoardInfo struct {
	BestUserName        User `json:"bestUserName"`
	BestUserOrdersCount int  `json:"bestUserOrdersCount"`
}
type Providers struct {
	Id       int    `json:"id"`
	Provider string `json:"provider"`
}
type Order struct {
	Id           int       `json:"id "`
	Date         time.Time `json:"date"`
	Price        int       `json:"price"`
	Cash         bool      `json:"cash"`
	ProductId    int       `json:"product"`
	ProductCount int       `json:"productCount"`
	ProductName  string    `json:"productName"`
	UserId       int       `json:"userId"`
	UserName     string    `json:"userName"`
	EmployeeId   int       `json:"employee"`
	EmployeeName string    `json:"employeeName"`
}

type Customer struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
}
