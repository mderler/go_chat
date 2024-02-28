package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mderler/go_chat/view/index"
)

type IndexHandler struct{}

func (h *IndexHandler) ShowIndex(c echo.Context) error {
	return render(c, index.Show())
}
