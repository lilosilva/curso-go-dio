package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Soma(a, b float64) float64 {
	return a + b
}

func Subtracao(a, b float64) float64 {
	return a - b
}

func Multiplicacao(a, b float64) float64 {
	return a * b
}

func Divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("erro: divisão por zero não é permitida")
	}
	return a / b, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Calculadora CLI em Go ---")

	for {
		fmt.Println("\n-----------------------------")
		fmt.Println("Escolha a operação (Soma, Subtracao, Multiplicacao, Divisao)")
		fmt.Print("Ou digite 'Sair' para terminar: ")

		operacaoInput, _ := reader.ReadString('\n')
		operacao := strings.TrimSpace(strings.ToLower(operacaoInput))

		if operacao == "sair" {
			fmt.Println("Encerrando a calculadora. Até mais!")
			break
		}

		if operacao != "soma" && operacao != "subtracao" && operacao != "multiplicacao" && operacao != "divisao" {
			fmt.Println("Operação inválida. Por favor, digite Soma, Subtracao, Multiplicacao ou Divisao.")
			continue
		}

		fmt.Print("Digite o primeiro número: ")
		num1Input, _ := reader.ReadString('\n')
		num1, err1 := parseInput(num1Input)
		if err1 != nil {
			fmt.Println("Entrada inválida para o primeiro número. Tente novamente.")
			continue
		}

		fmt.Print("Digite o segundo número: ")
		num2Input, _ := reader.ReadString('\n')
		num2, err2 := parseInput(num2Input)
		if err2 != nil {
			fmt.Println("Entrada inválida para o segundo número. Tente novamente.")
			continue
		}

		var resultado float64
		var err error

		switch operacao {
		case "soma":
			resultado = Soma(num1, num2)
			fmt.Printf("Resultado da Soma: %.2f\n", resultado)
		case "subtracao":
			resultado = Subtracao(num1, num2)
			fmt.Printf("Resultado da Subtração: %.2f\n", resultado)
		case "multiplicacao":
			resultado = Multiplicacao(num1, num2)
			fmt.Printf("Resultado da Multiplicação: %.2f\n", resultado)
		case "divisao":
			resultado, err = Divisao(num1, num2)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Resultado da Divisão: %.2f\n", resultado)
			}
		}
	}
}

func parseInput(input string) (float64, error) {

	cleanedInput := strings.TrimSpace(input)

	num, err := strconv.ParseFloat(cleanedInput, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
