package request

import "github.com/labstack/echo/v4"

type PostAccountSignup struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

type postAccountSignup struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

func NewPostAccountSignup(c echo.Context) (*PostAccountSignup, error) {
	o := new(postAccountSignup)
	if err := c.Bind(o); err != nil {
		return nil, err
	}

	return &PostAccountSignup{
		UserID:   o.UserID,
		Password: o.Password,
	}, nil
}
