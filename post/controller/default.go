package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetAll Get all posts
func GetAll(c echo.Context) error {
	return c.String(http.StatusOK, "Route GetAll posts")

}

// GetOne Get only post
func GetOne(c echo.Context) error {
	return c.String(http.StatusOK, "Route GetOne posts")
}

// Insert Insert post
func Insert(c echo.Context) error {
	return c.String(http.StatusOK, "Route Insert posts")
}

// Delete Delete post
func Delete(c echo.Context) error {
	return c.String(http.StatusOK, "Route Delete posts")
}

// Update Update post
func Update(c echo.Context) error {
	return c.String(http.StatusOK, "Route Update posts")
}
