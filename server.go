package main

import (
	"net/http"

	"bitbucket.org/eChurchIPB/api-echurch/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	routes.RegisterRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
