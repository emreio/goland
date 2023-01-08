package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	inmodulelib "mylearnings.go/main/InModuleLib"
	"mylearnings.go/main/business"
)

type ServerResponse struct {
	Name         string
	Surname      string
	Age          int
	IsRegistered bool
}

func main() {

	myHttpServer := inmodulelib.New("2222", "1")

	myHttpServer.AddHandler("/get", get)
	myHttpServer.AddHandler("/post", post)
	myHttpServer.AddHandler("/api/swagger", apiGet)
	myHttpServer.AddHandler("/api/query", apiQuery)
	myHttpServer.AddHandler("/getUsers", business.GetUsers)
	myHttpServer.AddHandler("/random", business.GetRandomNumber)

	myHttpServer.StartServer()
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Checking to see if Authorized header set...")

		if val, ok := r.Header["Authorized"]; ok {
			fmt.Println(val)
			if val[0] == "true" {
				fmt.Println("Header is set! We can serve content!")
				endpoint(w, r)
			}
		} else {
			fmt.Println("Not Authorized!!")
			fmt.Fprintf(w, "Not Authorized!!")
		}
	})
}

func aspectHandler(targetHandler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	fmt.Println("request recieved")
	fmt.Printf("id is %s")

	return targetHandler
}

func apiQuery(w http.ResponseWriter, r *http.Request) {

	urlValues := r.URL.Query()

	if urlValues["id"] != nil {
		fmt.Println("query not provided..")
	}

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

type myFunc func(string, string)

func (m myFunc) Serve() {

}
