package handler

import (
	"database/sql"
	"log"
	"strconv"

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
	userID := c.Get("userID").(int64)
	user, err := h.queries.GetUserForChatById(c.Request().Context(), userID)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return render(c, layout.ErrorBase(internalServerError))
	}

	contact, err := h.queries.GetLastContactedUser(c.Request().Context(), userID)
	if err != nil {
		log.Printf("Error getting last contacted user: %v", err)
		if err == sql.ErrNoRows {
			return render(c, index.Show(user, index.ChatParams{}, []model.GetContactedUsersRow{}))
		}
		return render(c, layout.ErrorBase(internalServerError))
	}

	messages, err := h.queries.GetDirectMessages(c.Request().Context(), model.GetDirectMessagesParams{
		UserID:    userID,
		ContactID: contact.ID,
	})
	if err != nil {
		log.Printf("Error getting messages: %v", err)
		return render(c, layout.ErrorBase(internalServerError))
	}

	contacedUsers, err := h.queries.GetContactedUsers(c.Request().Context(), userID)
	if err != nil {
		return render(c, layout.ErrorBase(internalServerError))
	}

	var messageViews []index.Message

	for _, message := range messages {
		messageViews = append(messageViews, index.Message{
			Author:  message.FullName,
			Color:   message.Color,
			Message: message.Content,
			Left:    message.SenderID != userID,
		})
	}

	chatParams := index.ChatParams{
		GroupName: contact.FullName,
		Color:     contact.Color,
		ContactID: contact.ID,
		Messages:  messageViews,
	}

	return render(c, index.Show(user, chatParams, contacedUsers))
}

func (h *IndexHandler) ShowChat(c echo.Context) error {
	userID := c.Get("userID").(int64)
	contactQuery := c.QueryParam("contact")

	contactID, err := strconv.ParseInt(contactQuery, 10, 64)
	if err != nil {
		return render(c, layout.ErrorBase(badRequest))
	}

	contact, err := h.queries.GetUserForChatById(c.Request().Context(), contactID)
	if err != nil {
		return render(c, layout.ErrorBase(internalServerError))
	}

	messages, err := h.queries.GetDirectMessages(c.Request().Context(), model.GetDirectMessagesParams{
		UserID:    userID,
		ContactID: contactID,
	})
	if err != nil {
		return render(c, layout.ErrorBase(internalServerError))
	}

	var messageViews []index.Message

	for _, message := range messages {
		messageViews = append(messageViews, index.Message{
			Author:  message.FullName,
			Color:   message.Color,
			Message: message.Content,
			Left:    message.SenderID != userID,
		})
	}

	chatParams := index.ChatParams{
		GroupName: contact.FullName,
		Color:     contact.Color,
		ContactID: contactID,
		Messages:  messageViews,
	}

	return render(c, index.ShowChat(chatParams))
}

func (h *IndexHandler) ShowContactedUsers(c echo.Context) error {
	userID := c.Get("userID").(int64)

	contacedUsers, err := h.queries.GetContactedUsers(c.Request().Context(), userID)
	if err != nil {
		return render(c, layout.ErrorBase(internalServerError))
	}

	return render(c, index.ShowContactedUsers(contacedUsers))
}
