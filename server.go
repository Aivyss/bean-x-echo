package main

import (
	_ "bean-x-echo/echo"
	_ "bean-x-echo/handler"
	_ "bean-x-echo/middleware"
	"github.com/aivyss/go-bean"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := bean.StartLazyLoading(); err != nil {
		panic(err)
	}

	e := bean.MustGetBean[*echo.Echo]()

	if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}
