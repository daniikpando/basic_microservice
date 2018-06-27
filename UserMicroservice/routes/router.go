package routes

import "github.com/labstack/echo"


// InitRouter se encarga de generar todas las rutas necesarias en la aplicación
func InitRouter()  *echo.Echo{

	e := echo.New()
	UserRouter(e)

	return e
}