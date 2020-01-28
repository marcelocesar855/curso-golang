package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	s2 := a2[1:3] // slice não é um array, e sim um pedaço de um array
	fmt.Println(a2, s2)

	s3 := a2[:2] // um novo slice, mas apontando para o mesmo array
	fmt.Println(a2, s3)

	s4 := s2[:1] //voce pode imaginar o slice como um tamanho definido com um ponteiro para um array
	fmt.Println(s2, s4)
}
