package wsserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pandoratoolbox/json"

	"github.com/go-chi/jwtauth"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type MessagePayload struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// Client represents a connected client
type Client struct {
	conn   *websocket.Conn
	userId int64
}

// ClientManager manages the connected clients
type ClientManager struct {
	clients    map[*Client]bool
	ClientById map[int64]*Client
	register   chan *Client
	unregister chan *Client
}

var Manager *ClientManager

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients:    make(map[*Client]bool),
		ClientById: make(map[int64]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (manager *ClientManager) Start() {
	for {
		select {
		case client := <-manager.register:
			manager.clients[client] = true
			manager.ClientById[client.userId] = client
			fmt.Println("Client Connected:", client.conn.RemoteAddr(), "Token:", client.userId)
		case client := <-manager.unregister:
			if _, ok := manager.clients[client]; ok {
				delete(manager.clients, client)
				delete(manager.ClientById, client.userId)
				client.conn.Close()
				fmt.Println("Client Disconnected:", client.conn.RemoteAddr(), "Token:", client.userId)
			}
		}
	}
}

func (manager *ClientManager) HandleConnection(w http.ResponseWriter, r *http.Request) {
	log.Println("1242413414n.")
	fmt.Println("------------")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		conn.Close()
		return
	}
	fmt.Println("------------1")

	accessToken := r.URL.Query().Get("token")
	if accessToken == "" {
		log.Println("No access token provided. Closing connection.")
		conn.Close()
		return
	}
	token, err := jwtauth.VerifyToken(jwtauth.New("HS256", []byte("h1l32b"), nil), accessToken)

	if err != nil {
		log.Println("Invalid token provided. Closing connection.")
		conn.Close()
		return
	}

	userId := token.PrivateClaims()["id"]

	client := &Client{conn: conn, userId: int64(userId.(float64))}
	manager.register <- client

	go manager.Read(client)
	// go manager.Write(client)
}

func (manager *ClientManager) Read(client *Client) {
	defer func() {
		manager.unregister <- client
		client.conn.Close()
	}()

	for {
		_, _, err := client.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func (manager *ClientManager) Write(client *Client, userid int64, msg MessagePayload) error {
	js, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return err
	}

	err = client.conn.WriteMessage(websocket.TextMessage, js)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("Message sent to client:", userid)

	return nil
}

func init() {
	Manager = NewClientManager()
	go Manager.Start()
}
