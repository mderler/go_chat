package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mderler/go_chat/view/index"
)

type IndexHandler struct{}

func (h *IndexHandler) ShowIndex(c echo.Context) error {
	messages := []index.Message{
		{Author: "Stefan Derler", Message: "Hello World!", Left: true},
		{Author: "Matthias Derler", Message: "Hello World!", Left: false},
	}
	return render(c, index.Show(messages))
}
