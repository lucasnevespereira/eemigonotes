package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.eemi.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			res, err := http.Get(url)
			if err != nil {
				log.Printf("Fetching %v: %v", url, err)
			}
			fmt.Println(res)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
