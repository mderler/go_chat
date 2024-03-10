package chat

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/binary"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mderler/go_chat/model"
	"github.com/mderler/go_chat/view/index"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024
)

type Message struct {
	isGroup       bool
	userOrGroupID int64
	userForChat   model.GetUserForChatByIdRow
	text          string
}

type Client struct {
	hub     *Hub
	db      *sql.DB
	queries *model.Queries
	userID  int64
	conn    *websocket.Conn
	send    chan Message
}

func unpackMessage(data []byte) (bool, int64, string, error) {
	if len(data) < 9 || len(data) > 9+1024 {
		return false, 0, "", errors.New("invalid message: insufficient data length")
	}

	messageType := data[0] != 0
	userOrGroupID := binary.LittleEndian.Uint64(data[1:9])
	text := string(data[9:])

	return messageType, int64(userOrGroupID), text, nil
}

func (c *Client) createMessageReturnSender(ctx context.Context, isGroup bool, userOrGroupID int64, text string) (Message, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return Message{}, err
	}
	defer tx.Rollback()

	qtx := c.queries.WithTx(tx)
	baseMessageID, err := qtx.CreateMessage(ctx, model.CreateMessageParams{SenderID: c.userID, Content: text})
	if err != nil {
		return Message{}, err
	}
	if !isGroup {
		err := qtx.CreateDirectMessage(ctx, model.CreateDirectMessageParams{MessageID: baseMessageID, ReceiverID: int64(userOrGroupID)})
		if err != nil {
			return Message{}, err
		}
	}

	userForChat, err := qtx.GetUserForChatById(ctx, c.userID)
	if err != nil {
		return Message{}, err
	}

	return Message{isGroup: isGroup, userOrGroupID: userOrGroupID, userForChat: userForChat, text: text}, tx.Commit()
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, rawMessage, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		isGroup, userOrGroupID, text, err := unpackMessage(rawMessage)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		message, err := c.createMessageReturnSender(context.TODO(), isGroup, userOrGroupID, text)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		c.hub.multicast <- message
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}

			var buf bytes.Buffer

			if message.isGroup {
				buf.WriteByte(1)
			} else {
				buf.WriteByte(0)
			}

			binary.Write(&buf, binary.LittleEndian, message.userOrGroupID)

			index.ChatMessage(message.userForChat.FullName, message.userForChat.Color, message.text, false).Render(context.TODO(), &buf)

			w.Write(buf.Bytes())

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func NewClient(hub *Hub, userID int64, queries *model.Queries, db *sql.DB, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, userID: userID, queries: queries, db: db, conn: conn, send: make(chan Message, 256)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
