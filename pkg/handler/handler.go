package handler

import (
	"course_work/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	test := router.Group("/test")
	{
		test.POST("/", h.test)
		test.GET("/", h.test)
		test.GET("/user", h.test)
		test.GET("/mod", h.test)
		test.GET("/admin", h.test)

	}

	api := router.Group("/api", h.userIdentity)
	{
		customer := api.Group("/customer")
		{
			customer.POST("/", h.CreateCustomer)
			customer.GET("/", h.GetCustomers)
			customer.GET("/:id", h.GetCustomers)
			customer.PATCH("/:id", h.UpdateCustomer)
			customer.DELETE("/:id", h.DeleteCustomer)

		}
		/*employee := api.Group("/employee")
		{
			employee.POST("/")
			employee.GET("/")
			employee.GET("/:id")
			employee.PUT("/:id")
			employee.DELETE("/:id")
		}
		item := api.Group("/item")
		{
			item.POST("/")
			item.GET("/")
			item.GET("/:id")
			item.PUT("/:id")
			item.DELETE("/:id")
		}
		itemInfo := api.Group("/itemInfo")
		{
			itemInfo.POST("/")
			itemInfo.GET("/")
			itemInfo.GET("/:id")
			itemInfo.PUT("/:id")
			itemInfo.DELETE("/:id")
		}
		mainStock := api.Group("/mainStock")
		{
			mainStock.POST("/")
			mainStock.GET("/")
			mainStock.GET("/:id")
			mainStock.PUT("/:id")
			mainStock.DELETE("/:id")
		}
		order := api.Group("/order")
		{
			order.POST("/")
			order.GET("/")
			order.GET("/:id")
			order.PUT("/:id")
			order.DELETE("/:id")
		}
		orderToStock := api.Group("/orderToStock")
		{
			orderToStock.POST("/")
			orderToStock.GET("/")
			orderToStock.GET("/:id")
			orderToStock.PUT("/:id")
			orderToStock.DELETE("/:id")
		}
		outlet := api.Group("/outlet")
		{
			outlet.POST("/")
			outlet.GET("/")
			outlet.GET("/:id")
			outlet.PUT("/:id")
			outlet.DELETE("/:id")
		}
		outletStock := api.Group("/outletStock")
		{
			outletStock.POST("/")
			outletStock.GET("/")
			outletStock.GET("/:id")
			outletStock.PUT("/:id")
			outletStock.DELETE("/:id")
		}
		provider := api.Group("/provider")
		{
			provider.POST("/")
			provider.GET("/")
			provider.GET("/:id")
			provider.PUT("/:id")
			provider.DELETE("/:id")
		}
		reciept := api.Group("/reciept")
		{
			reciept.POST("/")
			reciept.GET("/")
			reciept.GET("/:id")
			reciept.PUT("/:id")
			reciept.DELETE("/:id")
		}*/

	}
	return router
}
