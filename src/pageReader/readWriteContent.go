package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPage(url string) (int, error) {

	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()
	if body, err := ioutil.ReadAll(response.Body); err != nil {
		return 0, err
	} else {
		return len(body), nil
	}

}

func worker(urlCh chan string, sizeCh chan string) {
	url := <-urlCh
	if length, err := getPage(url); err == nil {
		sizeCh <- fmt.Sprintf("URL %s is %d long.", url, length)
	} else {
		sizeCh <- fmt.Sprintf("URL %s not found.", url)
	}

}

func generator(url string, urlCh chan string) {
	urlCh <- url
}
func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.yahoo.com",
		"http://www.bing.com",
		"http://www.amazon.com"}
	urlCh := make(chan string)
	pageLengthCh := make(chan string)

	for _, url := range urls {
		go worker(urlCh, pageLengthCh)
		go generator(url, urlCh)
	}
	for range urls {
		fmt.Printf("%s\n", <-pageLengthCh)
	}

}
