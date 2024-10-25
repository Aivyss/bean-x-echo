package main

import (
	_ "bean-x-echo/beans/echo"
	"bean-x-echo/beans/middleware"
	_ "bean-x-echo/handler"
	"github.com/aivyss/bean"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := bean.StartLazyLoading(); err != nil {
		panic(err)
	}

	e := bean.MustGetBean[*echo.Echo]()
	e.Use(bean.MustGetBean[*middleware.LogRequestMiddleware]().Process)

	if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}
