package main

import "fmt"

func main() {
	var name, lastName, favoriteFood string
	fmt.Println("Escribe tu nombre y apellido")
	fmt.Scanf("%s %s", &name, &lastName)
	fmt.Println("Escribe tu comida favorita")
	fmt.Scanf("%s", &favoriteFood)
	fmt.Println("Hi, my name is " + name + " " + lastName + " and my favorite food is " + favoriteFood)
}
