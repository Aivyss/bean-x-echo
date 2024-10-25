package handler

import (
	"bean-x-echo/dto/request"
	"bean-x-echo/usecase"
	"fmt"
	"github.com/aivyss/go-bean"
	"github.com/labstack/echo/v4"
)

func init() {
	bean.RegisterBeanLazy(NewAccountHandler)
}

type AccountHandler struct {
	accountUsecase usecase.AccountUsecase
}

func (a *AccountHandler) Signup(c echo.Context) error {
	req, err := request.NewPostAccountSignup(c)
	if err != nil {
		return err
	}

	return a.accountUsecase.Signup(c.Request().Context(), req.UserID, req.Password)
}

func NewAccountHandler(
	e *echo.Echo,
	accountUsecase usecase.AccountUsecase,
) *AccountHandler {
	fmt.Println("autowired: *AccountHandler")

	obj := &AccountHandler{
		accountUsecase: accountUsecase,
	}

	e.POST("/account/sign-up", obj.Signup)
	return obj
}
