package golangtryout

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Scanner is a adapter interface for scanning,
// could be website url, files, images and etc...
type Scanner interface {
	Scan()
}

type WebSite struct {
	Url string
}

// in this example, we only scan website
func (w *WebSite) Scan() {
	httpclient := &http.Client{}
	req, err := http.NewRequest("GET", w.Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := httpclient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func NewScanner(url string) Scanner {
	return &WebSite{Url: url}
}

func BatchScan(urls []string) {
	// implement goroutine to scan multiple urls
	// with number of workers defined
	// and wait for all workers to finish
	numWorkers := 5
	urlsChan := make(chan string, len(urls))
	doneChan := make(chan bool, numWorkers)
	for _, url := range urls {
		urlsChan <- url
	}
	//
	close(urlsChan)
	for i := 0; i < numWorkers; i++ {
		go func() {
			for url := range urlsChan {
				NewScanner(url).Scan()
			}
			doneChan <- true
		}()
	}
	for i := 0; i < numWorkers; i++ {
		<-doneChan
	}

}
