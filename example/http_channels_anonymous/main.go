package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkUrl(url string, c chan string) {

	resp, err := http.Get(url)
	if err != nil {
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error : %v\n", err)
		fmt.Println(s)
		c <- url
	} else {

		s := fmt.Sprintf("%s -> Status Code: %d\n", url, resp.StatusCode)

		s += fmt.Sprintf("%s is Up\n", url)
		fmt.Println(s)
		c <- url

	}

}

func main() {

	urls := []string{"https://www.golang.org", "https://www.google1.com", "https://www.medium.com"}

	c := make(chan string)

	for _, url := range urls {
		// the goroutine is sending the URL over channel for the next goroutine to use it.

		go checkUrl(url, c)
	}

	time.Sleep(time.Second * 2)

	// Spawning goroutines continously

	// 1.
	// for {
	//  go checkUrl(<-c, c)

	// OR

	// 2.
	// for url := range c { // It means "wait for the channel to return some values"
	//  fmt.Println(strings.Repeat("#", 30))
	//  time.Sleep(2 * time.Second) // pause for 2 seconds
	//  go checkUrl(url, c)
	// }

	// OR

	// 3.
	for url := range c {
		go func(u string) {
			time.Sleep(time.Second * 2)
			checkUrl(u, c)
		}(url)
	}
}
