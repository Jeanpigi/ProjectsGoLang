package main

import "fmt"

func main() {
	var edad int

	fmt.Println("Escribe tu edad")
	fmt.Scanf("%d", &edad)

	if edad >= 30 {
		fmt.Println("Nunca es tarde para aprender ¿Qué curso tomaremos?")
	} else if edad >= 18 && edad <= 29 {
		fmt.Println("Es un momento excelente para impulsar tu carrera")
	} else if edad >= 7 {
		fmt.Println("¿Sabes hacia dónde dirigir tu futuro? Seguro puedo ayudarte")
	}
}
