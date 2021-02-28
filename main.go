package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// New instance echo
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, This my first Demo CI/CD")
	})
	e.POST("/login", login)
	e.Logger.Fatal(e.Start(":8989"))
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	fmt.Println("Username: ", username)
	if username == "democircle" && password == "123456" {
		return c.JSON(http.StatusOK, map[string]string{
			"Message": "Login success",
		})
	}
	return c.JSON(http.StatusFound, map[string]string{
		"Message": "Login failed",
	})
}
