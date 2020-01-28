package main

import "fmt"

func notaPorConceito(n float64) string {
	if n > 9 && n < 10 {
		return "A"
	} else if n >= 8 && n <= 9 {
		return "B"
	} else if n >= 6 && n <= 8 {
		return "C"
	} else if n >= 4 && n <= 6 {
		return "D"
	} else if n >= 2 && n <= 4 {
		return "E"
	} else {
		return "F"
	}
}

func main() {
	fmt.Println(notaPorConceito(9.1))
	fmt.Println(notaPorConceito(7.5))
	fmt.Println(notaPorConceito(5.5))
	fmt.Println(notaPorConceito(2.1))
}
