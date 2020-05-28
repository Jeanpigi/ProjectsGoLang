package main

import "fmt"

func main() {
	var numeroLimite, numeroComparar int

	fmt.Println("Escriba el limite")
	fmt.Scanf("%d", &numeroLimite)
	fmt.Println("Escriba un número")
	fmt.Scanf("%d", &numeroComparar)

	if numeroLimite >= numeroComparar {
		fmt.Println("El número", numeroComparar, "se encuentra en el rango")
	} else {
		fmt.Println("El número", numeroComparar, "excede el limite permitido")
	}
}
