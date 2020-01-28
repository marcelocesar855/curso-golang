package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- fmt.Sprintf("%s falando: %d", pessoa, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case t1 := <-entrada1:
				ch <- t1
			case t2 := <-entrada2:
				ch <- t2
			}
		}
	}()
	return ch
}

func main() {
	ch := juntar(falar("Maria"), falar("João"))
	fmt.Println(<-ch, <-ch)
	fmt.Println(<-ch, <-ch)
	fmt.Println(<-ch, <-ch)
}
