package main

import "fmt"

func main() {
	var limiteSuperior, limiteInferior, numeroComparar int

	fmt.Println("Escriba el limite inferior")
	fmt.Scanf("%d", &limiteInferior)
	fmt.Println("Escriba el limite Superior")
	fmt.Scanf("%d", &limiteSuperior)
	fmt.Println("Escriba un número")
	fmt.Scanf("%d", &numeroComparar)

	if numeroComparar >= limiteInferior && numeroComparar <= limiteSuperior {
		fmt.Println("El número", numeroComparar, "se encuentra en el rango")
	} else {
		fmt.Println("El número", numeroComparar, "excede el limite permitido")
	}
}
