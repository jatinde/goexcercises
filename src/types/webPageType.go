package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type webPage struct {
	url  string
	body []byte
	err  error
}

func (w *webPage) get() {
	response, err := http.Get(w.url)
	if err != nil {
		w.err = err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		w.err = err
	}

	w.body = body
}
func main() {
	page := &webPage{url: "http://www.google.com"}
	page.get()
	if page.err != nil {
		fmt.Printf("URL: %s, Error: %s ", page.url, page.err.Error())
	} else {
		fmt.Printf("URL: %s, Content Length: %d ", page.url, len(page.body))
	}
}
