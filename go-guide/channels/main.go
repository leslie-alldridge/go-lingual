package main

import (
	"net/http"
	"fmt"
	// "os"
	// "io"
)

func main()  {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://trademe.co.nz",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err!= nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}