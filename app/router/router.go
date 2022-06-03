package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	handlef "HW/app/handler"
	"HW/app/middleware"
	"HW/app/repo"
	"HW/app/service"
	"HW/config"
	"HW/pkg/logger"
)

func NewRouter(handler *gin.Engine, l *logger.Logger, c *config.Config, db *gorm.DB) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Health probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Repo
	userRepo := repo.NewUserRepositoryInterface(db)
	categoryRepo := repo.NewCategoryRepository(db)
	productRepo := repo.NewProductRepository(db)
	cartRepo := repo.NewCartRepository(db)
	orderRepo := repo.NewOrderRepository(db)

	// Service
	authService := service.NewJWTAuthService(*c)
	userService := service.NewUserService(*userRepo)
	categoryService := service.NewCategoryService(*categoryRepo)
	productService := service.NewProductService(*productRepo)
	cartService := service.NewCartService(*cartRepo, *productRepo)
	orderService := service.NewOrderService(*orderRepo, *cartRepo)

	// Handler
	user := handlef.NewUser(*userService, *authService, *l)
	category := handlef.NewCategory(*categoryService, *l)
	product := handlef.NewProduct(*productService, *l)
	cart := handlef.NewCart(*cartService, *l)
	order := handlef.NewOrder(*orderService, *l)

	// Middleware
	authMw := middleware.NewJWTAuthMiddleware(*authService, *userService, *l)

	// Routers
	h := handler.Group("/api/v1")
	{
		a := h.Group("/user")
		{
			a.POST("/login", user.Login)
			a.POST("/signup", user.SignUp)
		}
		c := h.Group("/category", authMw.ValidateToken())
		{
			c.GET("", category.GetAllCategories)
			c.GET(":id", category.GetCategory)
			c.POST("", authMw.CheckRole("admin"), category.CreateCategory)
			c.DELETE(":id", authMw.CheckRole("admin"), category.DeleteCategory)
			c.POST("/bulk", authMw.CheckRole("admin"), category.CreateBulkCategory)
		}
		p := h.Group("/product", authMw.ValidateToken())
		{
			p.GET("", product.GetAllProducts)
			p.GET(":id", product.GetProduct)
			p.GET("/search/:query", product.SearchProducts)
			p.POST("", authMw.CheckRole("admin"), product.CreateProduct)
			p.PUT(":id", authMw.CheckRole("admin"), product.UpdateProduct)
			p.DELETE(":id", authMw.CheckRole("admin"), product.DeleteProduct)
		}
		b := h.Group("/cart", authMw.ValidateToken())
		{
			b.GET("", cart.GetCart)
			b.POST("", cart.AddCartItem)
			b.PUT("", cart.UpdateCartItem)
			b.DELETE("", cart.RemoveCartItem)
		}
		o := h.Group("/order", authMw.ValidateToken())
		{
			o.GET("", order.GetAllOrders)
			o.POST("", order.CreateOrder)
			o.PATCH(":id/cancel", order.CancelOrder)
		}
	}
}
