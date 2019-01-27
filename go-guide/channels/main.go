package main

import (
	"net/http"
	"fmt"
	"time"
	// "io"
)

func main()  {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://trademe.co.nz",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//normal loop

	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-c)
	// }

	//infinite loop

	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err!= nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}