package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		os.Exit(1)
	}
	// As Read inside the Reader interface can't resize the slice if it is full
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	// bs is itself updates as it is called by reference
	fmt.Println(string(bs))

}
