package channels

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetch(url string, out chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("error fetching %s: %v", url, err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	out <- fmt.Sprintf("%s: %d bytes", url, len(body))
}

func RunFinInOutParallelFetching() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://www.google.com",
	}

	results := make(chan string, len(urls))

	// fan-out: launch one goroutine per url
	for _, url := range urls {
		go fetch(url, results)
	}

	// fan-in: collect all results
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-results)
	}
}
