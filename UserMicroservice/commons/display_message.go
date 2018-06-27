package commons

import "github.com/labstack/echo"

func DisplayMessage(code int, message string, content interface{},context echo.Context) error{
	m := Message{code, message, content}
	return context.JSON(m.Code, m)
}