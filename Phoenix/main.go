package main

import "github.com/labstack/echo"

// La funcion principal
func main() {
	e := echo.New()

	e.Static("/" , "public")

	e.Start(":8080")
}
