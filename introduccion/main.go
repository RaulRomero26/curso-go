package main

import "fmt"

func main() {
	//declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)
	//declaracion de variables enteras
	base := 12 //los dos puntos indican que la variable no ha sido creada anteriormente
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado: ", areaCuadrado)
}
