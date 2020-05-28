package main

import "fmt"

//Sayaing especie en dragon ball
type Sayaing struct {
	Name  string
	Power int
}

//IntroduceGoku es la funcion que permite presentar a goku
func (s *Sayaing) IntroduceGoku() {
	fmt.Printf("Hi, I am %s. Ya! \n", s.Name)
}

//IntroduceVegeta es la funcion que permite presentar a vegeta
func (s *Sayaing) IntroduceVegeta() {
	fmt.Printf("My name is %s, Bastard \n", s.Name)
}

//SuperSayaing es la transformación de super sayaing
func (s *Sayaing) SuperSayaing() {
	s.Power += 100000
	fmt.Printf("My power is %v in supersayaing \n", s.Power)
}

func main() {
	goku := &Sayaing{
		Name:  "Goku",
		Power: 10000,
	}
	vegeta := &Sayaing{
		Name:  "Vegeta",
		Power: 9000,
	}

	goku.IntroduceGoku()
	goku.SuperSayaing()
	fmt.Println("Other Sayaing")
	vegeta.IntroduceVegeta()
	vegeta.SuperSayaing()
}
