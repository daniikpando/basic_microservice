package routes

import "github.com/labstack/echo"

func InitRouter()  *echo.Echo{

	e := echo.New()
	UserRouter(e)

	return e
}