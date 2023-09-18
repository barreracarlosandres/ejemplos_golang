package channel

import (
	"fmt"
	"net/http"
	"time"
)

var links = []string {
	"http://google.com",
	"http://facebook.com",
	"http://golang.org",
}

func Example1() {
	fmt.Println("------------ Channel Example 1 ------------")

	//Make a channel
	c := make(chan string)

	for _, link := range links {
		// go is a key work to get a new GO Routine (new thread)
		// Channel is used to comunnicate with the Go Routines
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		//recive the message of channel
		fmt.Println("executed:", <-c)
	}
}

//infinite loop of Go Routine with channel
func Example2() {
	fmt.Println("------------ Channel Example 2 ------------")
	
	//Make a channel
	c := make(chan string)

	for _, link := range links {
		// go is a key work to get a new GO Routine (new thread)
		// Channel is used to comunnicate with the Go Routines
		go checkLink(link, c)
	}

	//loop for each response
	/*for {
		//recive the message of channel
		go checkLinkLoop(<-c, c)
	}*/

	//loop for each responde, same to the for up
	for l := range c {
		//recive the message of channel
		//this is a anonimous funcion 
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!!!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}