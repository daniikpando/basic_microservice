package routes

import (
	"github.com/labstack/echo"
	"github.com/daniel/PruebaBackend/UserMicroservice/controllers/user"
)

func UserRouter(e *echo.Echo) {
	e.GET("/api/users", user.Index)
	e.POST("/api/users", user.Create)
}