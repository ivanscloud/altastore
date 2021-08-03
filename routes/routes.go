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

	// update profile
	jwtAuth.PUT("/customers/:id", controllers.UpdateProfileCustomersController)

	//Order Auth
	jwtAuth.POST("/orders", controllers.CreateOrders)
	jwtAuth.GET("/orders", controllers.GetOrder)

	// without jwt for login and register
	// route auth
	e.POST("/register", controllers.RegisterCustomersController)
	e.POST("/login", controllers.LoginCustomersController)
	return e
}
