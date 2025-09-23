package main

import "fmt"

var ebuicaoKelvin = 373.15

func main() {

	ebulicaoCelsius := ebuicaoKelvin - 273.15

	fmt.Printf("A temperatura de ebulição da água em Kelvin é: %.2f°F\n A temperatura de ebulição da água em Celsius é: %.2f°C\n", ebuicaoKelvin, ebulicaoCelsius)

}
