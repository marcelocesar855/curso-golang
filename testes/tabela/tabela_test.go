package strings

import (
	"strings"
	"testing"
)

const erroIndex = "%s (parte %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto string
		parte string
		index int
	}{
		{"Cod3r cursos online", "Cod3r", 0},
		{"", "", 0},
		{"Leonardo", "o", 2},
		{"Opa", "opa", -1},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)
		if atual != teste.index {
			t.Error(erroIndex, teste.texto, teste.parte, teste.index, atual)
		}
	}
}
