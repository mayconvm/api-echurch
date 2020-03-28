package routes

import (
	"bitbucket.org/eChurchIPB/api-echurch/post/controller"
	"github.com/labstack/echo"
)

// RoutesPost is to register all routes to Posts
func routesPost(e *echo.Echo) {
	e.GET("/posts", controller.GetAll)
	e.GET("/posts/:id", controller.GetOne)

	e.POST("/posts", controller.Insert)

	e.DELETE("/posts/:id", controller.Delete)
	e.PATCH("/posts/:id", controller.Update)
}
