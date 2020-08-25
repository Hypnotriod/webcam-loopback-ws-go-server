package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin:     checkOrigin,
}

func main() {
	startHTTPServer()
}

func startHTTPServer() {
	key := flag.String("key", "", "path to TLS key.pem")
	cert := flag.String("cert", "", "path to TLS cert.pem")
	address := flag.String("address", ":8080", "address, default is :8080")
	flag.Parse()

	http.Handle(
		"/",
		http.StripPrefix("/", http.FileServer(http.Dir("./public"))),
	)
	http.HandleFunc("/ws", serveWSRequest)
	if len(*key) > 0 && len(*cert) > 0 {
		http.ListenAndServeTLS(*address, "private/cert.pem", "private/key.pem", nil)
	} else {
		http.ListenAndServe(*address, nil)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func serveWSRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	log.Println("websocket open:", r.Host)
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func checkOrigin(r *http.Request) bool {
	return true
}