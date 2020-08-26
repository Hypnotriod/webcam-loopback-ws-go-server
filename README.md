# webcam-loopback-ws-go-server
Test webcam video loopback WebSocket server based on gorilla/websocket library written in `go`, configured for `VSCode`  
Check for the `-address` param in `lauch.json`  

### Installation

    go get github.com/gorilla/websocket

### Generation of cert.pem and key.pem

    openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem

## Lauch (Go plugin required)

`VSCode` -> `Ctrl + Shift + D`:  

* **Launch Debug** - to launch with HTTP (http://localhost:7777)  
* **Launch Debug TLS** - to launch with HTTPS (https://localhost:7777)   
