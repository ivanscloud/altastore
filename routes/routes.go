package routes

import (
	"altastore/constants"
	"altastore/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	jwtAuth := e.Group("")
	jwtAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	//Route Products
	jwtAuth.POST("/products", controllers.CreateProducts)
	jwtAuth.GET("/products", controllers.GetProducts)
	jwtAuth.PUT("/products/:id", controllers.UpdateProducts)
	jwtAuth.DELETE("/products/:id", controllers.DeleteProduct)
	// update profile
	jwtAuth.PUT("/customers/:id", controllers.UpdateProfileCustomersController)

	//Order Auth
	jwtAuth.POST("/orders", controllers.CreateOrders)
	jwtAuth.GET("/orders", controllers.GetOrder)

	// without jwt for login and register
	// route auth
	e.POST("/register", controllers.RegisterCustomersController)
	e.POST("/login", controllers.LoginCustomersController)

	//Route  CartItems
	jwtAuth.GET("/cartitems/:id", controllers.GetCartitemsByCartId)
	jwtAuth.POST("/cartitems", controllers.CreateCartitems)
	jwtAuth.PUT("/cartitems/:id", controllers.UpdateCartitems)
	jwtAuth.DELETE("/cartitems/:id", controllers.DeleteCartitems)

	//Route Carts
	jwtAuth.PUT("/carts/:id", controllers.UpdateCarts)

	return e
}
