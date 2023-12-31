package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {

	resp, err := http.Get(url)
	if err != nil {
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error : %v\n", err)
		c <- s //send error to channel
	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			file := strings.Split(url, "//")[1]
			file += ".txt"

			s += fmt.Sprintf("Writing response Body to %s\n", file)
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				// log.Fatal(err)
				s += "Error writing file\n"
				c <- s
			}
			s += fmt.Sprintf("%s is Up\n", url)
			c <- s
		}
	}

}

func main() {

	urls := []string{"https://www.golang.org", "https://www.google1.com", "https://www.medium.com"}

	c := make(chan string)

	for _, url := range urls {
		// 4.
		// Launch the goroutines
		go checkAndSaveBody(url, c)
	}

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	fmt.Println("No. of gorutines:", runtime.NumGoroutine())
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}
