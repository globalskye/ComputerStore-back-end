package handler

import (
	"course_work/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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

	product := router.Group("/product")
	{
		product.GET("/", h.GetAllProducts)
		product.GET("/categories", h.GetAllCategories)
		product.GET("/providers", h.GetAllProviders)
	}

	user := router.Group("/user", h.userIdentity)

	{

		card := user.Group("/card")
		{
			card.GET("/", h.GetUserCard)
			card.POST("/", h.PostToUserCard)
		}
		admin := user.Group("/admin", h.adminIdentity)
		{
			customer := admin.Group("/customer")
			{
				customer.GET("/", h.GetAllCustomers)
			}
			employee := admin.Group("/employee")
			{
				employee.GET("/", h.GetAllEmployee)
			}
			products := admin.Group("/product")
			{
				products.GET("/", h.GetAllProducts)
			}
			users := admin.Group("/users")
			{
				users.GET("/", h.GetAllUsers)
			}
			orders := admin.Group("/orders")
			{
				orders.GET("/", h.GetAllOrders)
			}
		}

	}

	return router
}
