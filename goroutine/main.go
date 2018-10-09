package main

import (
	"fmt"
	"log"
	"net/http"
)

// サイズがゼロの構造体
var empty struct{}

func getStatus(urls []string) <-chan string {
	statusChan := make(chan string, 3)
	// バッファを5にして生成
	limit := make(chan struct{}, 5)
	go func() {
		for _, url := range urls {
			select {
			case limit <- empty:
				go func(url string) {
					res, err := http.Get(url)
					if err != nil {
						log.Fatal(err)
					}
					defer res.Body.Close()
					statusChan <- res.Status
				}(url)
			}
		}
	}()
	return statusChan
}

func main() {
	// wait := new(sync.WaitGroup)
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
		"http://google.com",
	}

	statusChan := getStatus(urls)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
	// 待ち合わせる
	// wait.Wait()
}
