package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//This is a blocking operation, to read all the 5 channels we need to do the below
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// Instead we can do the below
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Repeating channels, this is an infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Refined way of writing infinite loop
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// Function Literal
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
