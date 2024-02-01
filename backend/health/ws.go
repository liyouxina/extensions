package main

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  8192,
	WriteBufferSize: 8192,
}

func wsServe() {
	http.HandleFunc("/health", wsHandler)
	err := http.ListenAndServeTLS(":10182", "./health.nxubiz.top.pem", "./health.nxubiz.top.key", nil)
	if err != nil {
		log.Error("接收ws连接出错 %s", err.Error())
	}

}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("接收连接出错", err.Error())
		return
	}
	defer conn.Close()
	people := &People{}
	for {
		_, messageByte, err := conn.ReadMessage()
		if err != nil {
			log.Error("ws读取数据出错", err.Error())
			break
		}
		message := string(messageByte)
		log.Info("Received message: %s", message)

		if strings.HasPrefix(message, "name=") {
			people.Name = strings.TrimLeft(message, "name=")
			peoplePool[people.Name] = people
		} else if strings.HasPrefix(message, "qrCode=") {
			people.QrCode = strings.TrimLeft(message, "qrCode=")
		}
	}
}
