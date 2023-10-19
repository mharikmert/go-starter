package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// define links
	main := "mharikmert.dev"
	subDomains := []string{"blog", "infra", "portal-backend", "portal-demo", "deployments"}

	// generate links
	links := []string{fmt.Sprintf("https://%s", main)}

	for _, subDomain := range subDomains {
		links = append(links, fmt.Sprintf("https://%s.%s", subDomain, main))
	}

	// // serial link checking
	// for _, link := range links {
	// 	checkLink(link) // blocking call (waits for each request to complete )
	// }

	// // link checking with go routine
	// for _, link := range links {
	// 	go checkLink(link) // go keyword launch a new routine
	// } // with this implementation(just using go keyword) main routine creates new child routines and continue and finish the main routine without waiting for the other ones.

	// create a channel for routines to communicate each other
	c := make(chan string)

	// link checking with go routine and channel
	for _, link := range links {
		go checkLinkWithChannel(link, c)
	}

	// // receive message from channel
	// fmt.Println(<-c) // this line of code will block the main routine until it receives a message from the channel

	// // receive message from channel in a loop
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// receive message from channel in a loop (alternative)
	for l := range c { // for loop will wait for a message from the channel and assign it to l

		// use function literal to launch a new routine without blocking the main routine
		go func(link string) {
			time.Sleep(5 * time.Second)      // sleep for 5 seconds
			go checkLinkWithChannel(link, c) // launch a new routine to check the link again
		}(l)
	}
}

func checkLink(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
	}

	fmt.Println(url, "is up!")
}

func checkLinkWithChannel(url string, c chan string) {
	_, err := http.Get(url) // blocking call
	// check if there is an error
	if err != nil {
		fmt.Printf("%s might be down!", url)
		c <- url // send message to channel
		return
	}
	fmt.Printf("%s is up!\n", url)
	c <- url // send message to channel
}
