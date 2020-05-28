package main

import "fmt"

func main() {
	var numero1, numero2 int
	fmt.Println("Escribe el primer número")
	fmt.Scanf("%d", &numero1)
	fmt.Println("Escribe el segundo número")
	fmt.Scanf("%d", &numero2)

	if numero1 > numero2 {
		fmt.Println("El número mayor es", numero1)
		fmt.Println("La diferencia es", numero1-numero2)
	} else if numero1 == numero2 {
		fmt.Println("Los números son iguales")
	} else {
		fmt.Println("El número mayor es", numero2)
		fmt.Println("la diferencia es", numero2-numero1)
	}
}
