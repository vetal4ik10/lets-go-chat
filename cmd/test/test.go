package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/vetal4ik10/lets-go-chat/configs"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}



var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


type Server struct {
	clients map[*client] bool
	send chan []byte
	clientConnect chan *client
}

func (s Server) run()  {
	for {
		select {
		case client := <-s.clientConnect:
			s.clients[client] = true
		case m := <-s.send:
			for c, a := range s.clients {
				if a {
					c.conn.WriteMessage(websocket.TextMessage, m)
				}
			}
		}

	}
}



type client struct {
	conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *client  {
	return &client{conn}
}

func (c *client) reader(s *Server)  {
	for {
		_, m, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		s.send <- m
	}
}


func (s *Server) WsConnect(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	c := NewClient(conn)
	s.clientConnect <- c

	go c.reader(s)
}



func main() {
	s := &Server{
		clients:   make(map[*client]bool),
		send: make(chan []byte),
		clientConnect: make(chan *client),
	}
	go s.run()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/ws", s.WsConnect)

	log.Fatal(http.ListenAndServe(":"+configs.GetServerPort(), r))



}
