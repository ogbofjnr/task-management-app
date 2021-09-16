package handlers

import (
	"github.com/ogbofjnr/maze/pkg/notificator"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type NotificationHandler struct {
	notificator *notificator.Notificator
	logger      *zap.Logger
}

func NewNotificationHandler(logger *zap.Logger, notificator *notificator.Notificator) *NotificationHandler {
	return &NotificationHandler{
		notificator: notificator,
		logger:      logger,
	}
}

// Connect serve websocket endpoint for push notifications
func (n *NotificationHandler) Connect(w http.ResponseWriter, r *http.Request) {
	conn, err := notificator.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &notificator.Client{Notificator: n.notificator, Conn: conn, Send: make(chan []byte, 256), UserID: 1}
	client.Notificator.Register <- client
	go client.WriteNotifications()
}
