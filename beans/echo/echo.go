package echo

import (
	"fmt"
	"github.com/aivyss/bean"
	"github.com/labstack/echo/v4"
)

func init() {
	fmt.Println("autowired: *echo.Echo")
	bean.RegisterBeanLazy(echo.New)
}
