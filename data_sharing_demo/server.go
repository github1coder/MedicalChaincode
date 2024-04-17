package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//升级为websocket协议
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client represents a connected client
type Client struct {
	conn *websocket.Conn
}

// Message represents a message from a client
type Message struct {
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	// Recipient string `json:"recipient"`
	Content []byte `json:"content"`
	DataID  string `json:"dataid"`
}

//连接客户端
var clients = make(map[*Client]bool)
//广播数据
var broadcast = make(chan Message)

//建立连接
func handleConnections(w http.ResponseWriter, r *http.Request) {
	//升级协议
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	//新建客户端连接
	client := &Client{conn: conn}
	clients[client] = true

	fmt.Println("Client connected")

	//监听读取数据
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			delete(clients, client)
			break
		}

		//广播数据
		broadcast <- msg
	}
}

//处理客户端消息
func handleMessages() {
	for {
		msg := <-broadcast

		//遍历所有连接客户端，广播数据
		for client := range clients {
			if client.conn != nil {
				//发送给当前客户端
				err := client.conn.WriteJSON(msg)
				if err != nil {
					log.Println(err)
					client.conn.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func main() {
	http.HandleFunc("/", handleConnections)
	go handleMessages()

	fmt.Println("WebSocket server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
