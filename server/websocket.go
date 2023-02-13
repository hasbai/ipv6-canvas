package server

import (
	"github.com/hasbai/ipv6-canvas/lib"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"strings"
)

type Hub struct {
	clients     map[*Client]bool
	clientCount int
	broadcast   chan []byte
	register    chan *Client
	unregister  chan *Client
}

var hub = Hub{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.clientCount++
			log.Printf("Registed %v, Client count: %d", client, h.clientCount)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				h.clientCount--
			}
			log.Printf("Unregisted %v, Client count: %d", client, h.clientCount)
		case message := <-h.broadcast:
			for client := range h.clients {
				client.send <- message
			}
		}
	}
}

type Client struct {
	conn  *websocket.Conn
	send  chan []byte
	close chan int
}

func (c *Client) Read() {
	for {
		data := make([]byte, 64)
		n, err := c.conn.Read(data)
		if err != nil {
			c.Close()
			return
		}
		var p lib.Pixel
		err = p.Unmarshal(data[:n])
		if err != nil {
			c.Close(err.Error())
			return
		}
		modifyImage(p)
	}
}

func (c *Client) Write() {
	for {
		select {
		case message := <-c.send:
			_, err := c.conn.Write(message)
			if err != nil {
				log.Println(err)
				c.Close()
				return
			}
		case <-c.close:
			return
		}
	}
}

func (c *Client) Close(message ...string) {
	if len(message) > 0 {
		c.conn.Write([]byte(strings.Join(message, " ")))
	}
	c.close <- 0
	hub.unregister <- c
}

func modifyImage(p lib.Pixel) {
	IMG.Modify(p)
	hub.broadcast <- p.Marshal()
}

func handleWS(ws *websocket.Conn) {
	client := Client{
		conn:  ws,
		send:  make(chan []byte),
		close: make(chan int),
	}
	go client.Read()
	go client.Write()
	client.send <- []byte("Hello")
	hub.register <- &client
	code := <-client.close
	log.Printf("Client closed with code %d", code)
}

func init() {
	go hub.Run()
	http.Handle("/ws", websocket.Server{
		Handler: handleWS,
	})
}
