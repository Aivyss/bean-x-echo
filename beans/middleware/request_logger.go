package middleware

import (
	"fmt"
	"github.com/aivyss/bean"
	"github.com/labstack/echo/v4"
	"io"
)

func init() {
	bean.RegisterBeansLazy(NewLogRequestMiddleware)
}

type LogRequestMiddleware struct{}

func NewLogRequestMiddleware() *LogRequestMiddleware {
	fmt.Println("autowired: *LogRequestMiddleware")
	return &LogRequestMiddleware{}
}

func (m *LogRequestMiddleware) Process(a echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Path :", c.Request().URL.Path)
		fmt.Println("Method :", c.Request().Method)
		body, err := io.ReadAll(c.Request().Body)
		if err == nil {
			fmt.Println("Body :", string(body))
		}

		return a(c)
	}
}
