package routes

import (
	"github.com/labstack/echo"
)

// RegisterRoutes is to register all routes
//TODO https://echo.labstack.com/guide/routing
func RegisterRoutes(e *echo.Echo) {

	//posts
	routesPost(e)
}
