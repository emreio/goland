package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	inmodulelib "mylearnings.go/main/InModuleLib"
)

type ServerResponse struct {
	Name         string
	Surname      string
	Age          int
	IsRegistered bool
}

func main() {
	fmt.Println("emre kantar")

	myHttpServer := inmodulelib.New("2222", "1")

	myHttpServer.AddHandler("/get", get)
	myHttpServer.AddHandler("/post", post)
	myHttpServer.AddHandler("/api/get", apiGet)

	myHttpServer.StartServer()
}

func apiGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(403)
		w.Write([]byte("METHOD NOT ALLOWED"))
		return
	}

	res, err := json.Marshal(ServerResponse{Name: "Emre", Surname: "Kantar", Age: 35, IsRegistered: true})

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("my lovely post handler"))
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello go"))
}

func GetSomeData(obj *interface{}) *interface{} {
	return obj
}
