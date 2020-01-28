package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2() string {
	return "F2"
}

func f3(p1 string, p2 string) {
	fmt.Printf("F3: %s %s\n", p1, p2)
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5(p1, p2 string) (string, string) { //em go existe retornos multiplos
	return p1, p2
}

func main() {
	f1()
	fmt.Println(f2())
	f3("Param1", "Param2")
	fmt.Println(f4("Param1", "Param2"))

	r1, r2 := f5("Param1", "Param2")
	fmt.Printf("F5: %s %s", r1, r2)
}
