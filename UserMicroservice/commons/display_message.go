package commons

import "github.com/labstack/echo"


// DisplayMessage se encarga de enviar la respuesta al cliente
func DisplayMessage(code int, message string, content interface{},context echo.Context) error{
	m := Message{code, message, content}
	return context.JSON(m.Code, m)
}