package main

import "fmt"

type animals interface {
	move() string
}

type dog struct {
	name        string
	description string
}

type cat struct {
	name        string
	description string
}

func (*dog) move() string {
	return "The dog is moving"
}

func (*cat) move() string {
	return "The cat is moving"
}

func animalMove(a animals) {
	fmt.Println(a.move())
}

func main() {
	momina := &cat{
		name:        "Momina",
		description: "I am an old Cat",
	}
	gata := &cat{
		name:        "Gata",
		description: "I am a crazy cat",
	}
	beckham := &dog{
		name:        "Beckham",
		description: "I am a black dog",
	}

	fmt.Printf("my name is %s. %s \n", beckham.name, beckham.description)
	fmt.Printf("my name is %s. %s \n", momina.name, momina.description)
	fmt.Printf("My name is %s. %s \n", gata.name, gata.description)
	animalMove(beckham)
	animalMove(momina)
}
