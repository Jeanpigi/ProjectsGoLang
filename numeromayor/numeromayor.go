package main

import "fmt"

func main() {
	var primerNumero, segundoNumero, tercerNumero int
	fmt.Println("Escriba el primer número")
	fmt.Scanf("%d", &primerNumero)
	fmt.Println("Escriba el segundo número")
	fmt.Scanf("%d", &segundoNumero)
	fmt.Println("Por ultimo escriba el tercer número")
	fmt.Scanf("%d", &tercerNumero)

	if primerNumero > segundoNumero {
		if primerNumero > tercerNumero {
			fmt.Println("El número mayor es", primerNumero)
		} else {
			fmt.Println("El número mayor es", tercerNumero)
		}
	} else {
		if segundoNumero > tercerNumero {
			fmt.Println("El número mayor es", segundoNumero)
		} else {
			fmt.Println("El número mayor es", tercerNumero)
		}
	}
}
