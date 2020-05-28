package main

import "fmt"

func main() {
	var isRaining string
	var isWindy string

	fmt.Println("Esta lloviendo")
	fmt.Scanf("%s", &isRaining)

	if isRaining == "si" || isRaining == "Si" || isRaining == "SI" {
		fmt.Println("Está haciendo mucho viento")
		fmt.Scanf("%s", &isWindy)
		if isWindy == "si" || isWindy == "Si" || isWindy == "SI" {
			fmt.Println("Hace mucho viento para salir con sombrilla")
		} else {
			fmt.Println("Deberias llevar una sombrilla")
		}
	} else {
		fmt.Println("Que tengas un buen día")
	}
}
