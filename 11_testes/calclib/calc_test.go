package calclib

import "testing"

func TestSoma(t *testing.T) {
	res := Soma(1, 3)

	if res != 4 {
		t.Error("Experado 4, mas o resultado foi ", res)
	}
}

func TestMultiplica(t *testing.T) {
	res := Multiplica(1, 3)

	if res != 3 {
		t.Error("Experado 3, mas o resultado foi ", res)
	}
}
