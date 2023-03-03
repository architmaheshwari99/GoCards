package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		os.Exit(1)
	}
	// // As Read inside the Reader interface can't resize the slice if it is full
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// // bs is itself updates as it is called by reference
	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	return len(b), nil
}
