package notificator

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"time"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Notifiable interface {
	GetMessage() ([]byte, error)
}

type Notificator struct {
	Clients       map[*Client]bool
	notifications chan *Notification
	Register      chan *Client
	Unregister    chan *Client
	Logger        *zap.Logger
}

type Notification struct {
	UserID  int
	Message []byte
}

func NewNotificator(logger *zap.Logger) *Notificator {
	return &Notificator{
		notifications: make(chan *Notification),
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
		Clients:       make(map[*Client]bool),
	}
}

func (h *Notificator) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case notification := <-h.notifications:
			connections := h.findUserConnections(notification.UserID)
			for _, client := range connections {
				select {
				case client.Send <- notification.Message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

func (h *Notificator) Notify(userID int, notification Notifiable) {
	message, err := notification.GetMessage()
	if err != nil {
		h.Logger.Error(err.Error())
	}
	h.notifications <- &Notification{userID, message}
}

func (h *Notificator) findUserConnections(userID int) []*Client {
	var userClients []*Client
	for c := range h.Clients {
		if c.UserID == userID {
			userClients = append(userClients, c)
		}
	}
	return userClients
}
