package matematica

import "testing"

const erroPadrao = "Valor esperado: %v\nValor obtido: %v"

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valorEsperado != valor {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
