package main

import (
	"fmt"
	"net/http"

	inmodulelib "mylearnings.go/main/InModuleLib"
)

func main() {
	fmt.Println("emre kantar")

	myHttpServer := inmodulelib.New("2222", "1")

	myHttpServer.AddHandler("/get", get)
	myHttpServer.AddHandler("/post", post)

	//myHttpServer.StartServer()
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
