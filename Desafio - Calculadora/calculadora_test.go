package main

import "testing"

func TestShouldSomaCorrect(t *testing.T) {

	esperado := 15.5
	resultado := Soma(10.5, 5.0)

	if resultado != esperado {
		t.Errorf("Esperado: %.2f, Resultado: %.2f", esperado, resultado)
	}

}

func TestShouldSubtracaoCorrect(t *testing.T) {

	esperado := 12.5
	resultado := Subtracao(20.0, 7.5)

	if resultado != esperado {
		t.Errorf("Esperado: %.2f, Recebido: %.2f", esperado, resultado)
	}

}

func TestShouldMultiplicacaoCorrect(t *testing.T) {

	esperado := 4.0
	resultado := Multiplicacao(8.0, 0.5)

	if resultado != esperado {
		t.Errorf("Esperado: %.2f, Recebido: %.2f", esperado, resultado)
	}

}

func TestShouldDivisaoCorrect(t *testing.T) {

	esperado := 5.0
	resultado, err := Divisao(15.0, 3.0)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if resultado != esperado {
		t.Errorf("Esperado: %.2f, Recebido: %.2f", esperado, resultado)
	}

}
