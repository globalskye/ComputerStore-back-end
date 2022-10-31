package handler

import (
	"course_work/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	/*
		api := router.Group("/api")
		{
			customer := api.Group("/customer")
			{
				customer.POST("/")
				customer.GET("/")
				customer.GET("/:id")
				customer.PUT("/:id")
				customer.DELETE("/:id")

			}
			employee := api.Group("/employee")
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
			}

		}*/
	return router
}
