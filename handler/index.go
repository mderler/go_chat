package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mderler/go_chat/model"
	"github.com/mderler/go_chat/view/index"
	"github.com/mderler/go_chat/view/layout"
)

type IndexHandler struct {
	queries *model.Queries
}

func NewIndexHandler(queries *model.Queries) *IndexHandler {
	return &IndexHandler{queries: queries}
}

func (h *IndexHandler) ShowIndex(c echo.Context) error {
	messages := []index.Message{
		{Author: "Stefan Derler", Message: "Hello World!", Left: true},
		{Author: "Matthias Derler", Message: "Hello World!", Left: false},
	}

	userID := c.Get("userID").(int64)
	fullname, err := h.queries.GetFullNameById(c.Request().Context(), userID)
	if err != nil {
		return render(c, layout.ErrorBase(internalServerError))
	}

	return render(c, index.Show(fullname, messages))
}
