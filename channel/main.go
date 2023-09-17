package channel

import (
	"fmt"
	"net/http"
)


func Example1() {
	fmt.Println("------------ Channel Example 1 ------------")
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
	}

	//Make a channel
	c := make(chan string)

	for _, link := range links {
		// go is a key work to get a new GO Routine (new thread)
		// Channel is used to comunnicate with the Go Routines
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		//recive the message of channel
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!!!")
		//send message to channel
		c <- "Might be down I think"
		return
	}


	fmt.Println(link, "is up!")
	c <- "Yeap its up"
}