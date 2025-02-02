package main

import (
	"fmt"
	"net/http"
	"sync"
)


func returneType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	fmt.Printf("%s: %s\n", url, ctype)
}


func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string){
			returneType(url)
			wg.Done()
			}(url)
	}

	wg.Wait()

}