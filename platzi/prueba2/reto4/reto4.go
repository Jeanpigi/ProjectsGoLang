package main

import "fmt"

func main() {
	var animalFavorito string

	fmt.Println("Escribe tu animal favorito")
	fmt.Scanf("%s", &animalFavorito)

	if animalFavorito == "tortuga" || animalFavorito == "Tortuga" || animalFavorito == "TORTUGA" || animalFavorito == "tortugas" || animalFavorito == "Tortugas" || animalFavorito == "TORTUGAS" {
		fmt.Println("También me gusta las tortugas")
	} else {
		fmt.Println("Ese animal es genial, pero prefiero las tortugas")
	}
}
