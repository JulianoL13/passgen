package keygen

import (
	"testing"
)

func TestSizeGetRandomPass(t *testing.T) {
	got := GetRandomPass(10, true, true, true)
	want := 10

	if len(got) != want {
		t.Errorf("O valor esperado era %d, foi encontrado: %d", want, len(got))
	}

}

func TestRandomGetRandomPass(t *testing.T) {
	got1 := GetRandomPass(10, true, true, true)
	got2 := GetRandomPass(10, true, true, true)
	got3 := GetRandomPass(10, true, true, true)

	if got1 == got2 || got1 == got3 || got2 == got3 {
		t.Errorf("Os valores não podem ser iguais, %s, %s, %s", got1, got2, got3)
	}
}

func TestAddRequiredChar(t *testing.T) {
	test := "ola"
	addRequiredChar(true, &test, "joão")
	got := test
	want := "olajoão"

	if got != want {
		t.Errorf("O valor esperado era %s, foi encontrado %s", want, got)
	}
}
