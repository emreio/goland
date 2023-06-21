package main

import (
	"fmt"
	"os"
)

func main() {

	_emre := os.Getenv("emre")

	for _, e := range os.Environ() {
		println(e)
	}

	println(_emre)

	user := User{mail: "emre@kantar.com", username: "kantar", age: 34}
	user.buffer = []byte("ewr")

	println("users initial email :" + user.mail)
	fmt.Printf("pointer is %p \r\n", &user)

	user.changeMe("emrekantar@gmail.com")
	// fmt.Println("user email changed to :" + user.mail)

	user.changeMeByFunc("emrekantar@gmail.com")
	//fmt.Println("user email changed to :" + user.mail)

	fmt.Println("go routine finished..")

	a := "emrğç"

	buffer := []byte(a)

	for _, r := range a {
		println(r)
	}

	l := len(buffer)

	println(l)
}

func (user User) changeMe(email string) {
	fmt.Printf("pointer is %p \r\n", &user)
	user.mail = email
}

func (user *User) changeMeByFunc(email string) {
	fmt.Printf("pointer is %p \r\n", user)
	user.mail = email
}

type User struct {
	mail     string
	username string
	age      int
	buffer   []byte
}
