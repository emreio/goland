package business

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
)

func GetRandomNumber(w http.ResponseWriter, r *http.Request) {

	var randomBuffer [20]byte
	io.ReadFull(rand.Reader, randomBuffer[:])

	w.Header().Set("Content-Type", "text/plain")
	w.Write("random number buffer")

	fmt.Println("random bytes sent to client")
}
