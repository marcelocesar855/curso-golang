package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("O literal inteiro é", reflect.TypeOf(3200))

	//sem sinal (naturais) uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O tipo sem sinal é", reflect.TypeOf(b))

	//com sinal (inteiros) int8 int16 int32 int64
	i := math.MaxInt64
	fmt.Println("O valor máximo de um inteiro é", i)
	fmt.Println("O tipo do inteiro é", reflect.TypeOf(i))

	var i1 rune = 'a' //rune representa o valor da string na tabela unicode
	fmt.Println("O tipo rune é do tipo", reflect.TypeOf(i1))
	fmt.Println("O valor de a na tabela unicode é", i1)

	//numeros com ponto flutuante (reais) (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo literal de um float é", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s := "Olá meu nome é Marcelo"
	fmt.Println(s + "!")
	fmt.Println(len(s))

	//string multiplas linhas
	s1 := `	Ola
	meu
	nome
	é
	Marcelo`
	fmt.Println(s1)

	//char?
	//var x1 char = "a" - tipo char não existe em go

	x1 := 'a' //literalmente um rune
	fmt.Println("O tipo de char é", reflect.TypeOf(x1))
	fmt.Println(x1)
}
