package main

import "math"

// Inicializando uma variável ou função ou qualquer tipo de objeto com letra maiúscula significa que é PÚBLICO (visível no pacote e fora dele)
// Inicializando com letra minúscula significa que é PRIVADO (visível apenas no pacote)
// Em Go não existe visíbilidade por arquivo, e sim por pacote, não podendo haver objetos duplicados nos pacotes

// Por exemplo...
// Ponto gerará algo público
// ponto ou _Ponto gerará algo privado

//Objetos públicos necessitam de comentários explicando o objeto (algo relevante)

//Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(a.x - a.y)
	cy = math.Abs(b.x - b.y)
	return
}

//Distancia calcula a distancia entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
