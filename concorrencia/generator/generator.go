package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//<- chan : somente canal de leitura

func titulo(urls ...string) <-chan string { //encapsula a criação do canal e da goroutine
	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title[^>]*>([^<]+)<\\/title>")
			ch <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return ch
}

func main() {
	t1 := titulo("https://www.google.com", "https://www.udemy.com")   //consome o canal gerado
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com") //consome o canal gerado
	fmt.Println("Primeiros: ", <-t1, " | ", <-t2)
	fmt.Println("Segundos: ", <-t1, " | ", <-t2)
}
