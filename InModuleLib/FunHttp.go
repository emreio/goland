package inmodulelib

import (
	"fmt"
	"log"
	"net/http"
)

type MyHttpServer struct {
	port    string
	timeout string
	handler map[string]func(w http.ResponseWriter, r *http.Request)
}

func New(port string, timeout string) *MyHttpServer {
	var instance = &MyHttpServer{port: port, timeout: timeout}
	instance.handler = make(map[string]func(w http.ResponseWriter, r *http.Request))
	return instance
}

func (httpServer *MyHttpServer) AddHandler(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	httpServer.handler[path] = handlerFunc
}

func (httpServer *MyHttpServer) StartServer() {

	fmt.Println("Registering http handlers.")
	var count int

	for key, value := range httpServer.handler {
		http.HandleFunc(key, value)
		count++
	}

	fmt.Printf("Total handler count: %v\n", count)

	fmt.Println("Server Started Listening At Port: " + httpServer.port)

	log.Fatal(http.ListenAndServe(":"+httpServer.port, nil))
}
