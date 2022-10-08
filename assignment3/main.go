package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Element struct {
	Wind        int
	Water       int
	SendMessage bool
}

func NewElement() *Element {
	return &Element{
		SendMessage: true,
	}
}

func (e *Element) rerollElement() {
	rand.Seed(time.Now().UnixMicro())
	for {
		e.Water = rand.Intn(100)
		e.Wind = rand.Intn(100)
		e.SendMessage = true
		time.Sleep(15 * time.Second)
	}
}

type webSocket struct {
	Element *Element
}

func NewWebSocket(e *Element) *webSocket {
	return &webSocket{
		Element: e,
	}
}

func (w *webSocket) webSocketService(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	payload := map[string]interface{}{
		"wind":  &w.Element.Wind,
		"water": &w.Element.Water,
	}

	for {
		for w.Element.SendMessage == true {
			if err := ws.WriteJSON(payload); err != nil {
				panic(err)
			}
			w.Element.SendMessage = false
		}
	}
}

func main() {
	element := NewElement()
	go element.rerollElement()

	ws := NewWebSocket(element)

	app := gin.Default()
	app.GET("/websocket", ws.webSocketService)
	app.GET("/", func(c *gin.Context) {
		c.File("./views/index.html")
	})
	app.StaticFile("/styles.css", "./views/styles.css")

	app.Run(":8080")
}
