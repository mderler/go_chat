package handler

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mderler/go_chat/model"
	"github.com/mderler/go_chat/view/user"
)

type UserHandler struct {
	queries *model.Queries
}

func NewUserHandler(queries *model.Queries) *UserHandler {
	return &UserHandler{queries: queries}
}

func (h *UserHandler) ShowUserList(c echo.Context) error {
	query := fmt.Sprintf("%%%s%%", c.QueryParam("q"))

	users, err := h.queries.GetUsersByQuery(c.Request().Context(), query)
	if err != nil {
		log.Println(err)
		return err
	}

	return render(c, user.ShowUserList(users))
}
