package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		c <- "Might be down"
	} else {
		fmt.Println("Website OK")
		c <- "It's up"
	}
	c <- link

}

func main() {
	links := []string{"https://pkg.go.dev/io#Copy", "https://music.youtube.com/", "https://www.google.com/"}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)

	}

	for {
		go checkLink(<-c, c)
	}

}
