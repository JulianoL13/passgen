package keygen

import (
	"testing"
)

func TestSizeGetRandomPass(t *testing.T) {
	got, err := GetRandomPass(10, true, true, true)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	want := 10
	if len(got) != want {
		t.Errorf("O valor esperado era %d, foi encontrado: %d", want, len(got))
	}
}

func TestRandomGetRandomPass(t *testing.T) {
	got1, err1 := GetRandomPass(10, true, true, true)
	if err1 != nil {
		t.Errorf("Erro inesperado: %v", err1)
	}
	got2, err2 := GetRandomPass(10, true, true, true)
	if err2 != nil {
		t.Errorf("Erro inesperado: %v", err2)
	}
	got3, err3 := GetRandomPass(10, true, true, true)
	if err3 != nil {
		t.Errorf("Erro inesperado: %v", err3)
	}

	if got1 == got2 || got1 == got3 || got2 == got3 {
		t.Errorf("Os valores não podem ser iguais, %s, %s, %s", got1, got2, got3)
	}
}

func TestAddRequiredChar(t *testing.T) {
	test := "ola"
	err := addRequiredChar(true, &test, "joão")
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	got := test
	want := "olajoão"

	if got != want {
		t.Errorf("O valor esperado era %s, foi encontrado %s", want, got)
	}

	test = "ola"
	err = addRequiredChar(false, &test, "joão")
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	got = test
	want = "ola"
	if got != want {
		t.Errorf("O valor esperado era %s, foi encontrado %s", want, got)
	}
}
