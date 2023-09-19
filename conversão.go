package main

import "fmt"

const ebulicaoK = 373.0

func main() {

	tempK := ebulicaoK
	celsius := tempK - 273

	fmt.Printf("O ponto de ebulição da água em K é %g  e em °C é %g ", ebulicaoK, celsius)
}
