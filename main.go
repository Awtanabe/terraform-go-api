// main.go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from Echo!")
	})
	// 再度テスト
	// テスト
	// 再度テスト
	// 再度テスト
	e.Logger.Fatal(e.Start(":8080"))
}
