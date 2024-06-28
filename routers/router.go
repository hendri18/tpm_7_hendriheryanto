package routers

import (
	"tpm_7_HendriHeryanto/handler"
	"tpm_7_HendriHeryanto/middleware"
	"tpm_7_HendriHeryanto/repository"
	"tpm_7_HendriHeryanto/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := &repository.UserRepo{DB: db}
	userService := &service.UserService{UserRepo: userRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	userRouter := router.Group("/users")
	userRouter.POST("/login", userHandler.Login)
	userRouter.POST("/register", userHandler.Register)

	productRepo := &repository.ProductRepo{DB: db}
	productService := &service.ProductService{ProductRepo: productRepo}
	productHandler := &handler.ProductHandler{ProductService: productService}

	productRouter := router.Group("/products")
	productRouter.Use(middleware.Authentication())
	productRouter.GET("", productHandler.Get)
	productRouter.POST("", productHandler.Create)
	productRouter.Use(middleware.ProductAuthorization(productService))
	productRouter.PUT("/:id", productHandler.Update)
	productRouter.DELETE("/:id", productHandler.Delete)

	return router
}
