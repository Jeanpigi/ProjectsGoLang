package main

import "fmt"

func main() {
	var creditos, semestre int
	fmt.Println("Escribir el número de créditos")
	fmt.Scanf("%d", &creditos)
	if creditos >= 1 && creditos <= 150 {
		semestre = creditos / 15
		if semestre > 0 {
			fmt.Printf("El estudiante se encuentra cursando el %v semestre. \n", semestre)
		} else {
			fmt.Println("El estudiante se encuentra cursando el 1 semestre")
		}
	} else {
		fmt.Println("No puede tener créditos menores de cero")
	}
}
