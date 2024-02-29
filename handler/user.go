package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mderler/go_chat/view/user"
)

type UserHandler struct{}

func (h *UserHandler) ShowUser(c echo.Context) error {
	return render(c, user.Show())
}
