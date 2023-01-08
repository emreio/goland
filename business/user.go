package business

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	Name        string
	Surname     string
	DateofBirth string
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("incoming getUsers request : %s | %s \n", r.RemoteAddr, r.UserAgent())

	file, err := os.OpenFile("./data.dat", os.O_RDONLY, 0644)

	if err != nil {
		fmt.Errorf("error occured while opening data.dat file : %e", err)
		w.WriteHeader(500)
		w.Write([]byte("an internal error occured"))
		return
	}

	defer file.Close()

	buff, err := io.ReadAll(file)

	if err != nil {
		fmt.Errorf("error occured %e", err)
		w.WriteHeader(500)
		w.Write([]byte("an internal error occured"))
		return
	} else {
		w.Write(buff)
	}
}
