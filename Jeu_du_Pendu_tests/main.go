package main

import (
	"fmt"

	"github.com/Go-tuto/Jeu_du_Pendu_tests/player"
)

func main() {
	var p1 player.Gamer
	p1.Name = "Maka"
	p1.Age = "24"

	fmt.Printf("Gamer 1 %v\n", p1)
	fmt.Printf("p1.Name=%v, p1.Age=%v\n", p1.Name, p1.Age)

	a := player.Avatar{"http://avatar.jpg"}
	fmt.Printf("avatar : %v\n", p1, a)

	p3 := player.Gamer{
		Name: "MAka",
		Age:  25,
		Avatar: player.Avatar{
			Url: "http://url.com",
		},
	}
	fmt.Printf("Gamer 3 : %v\n", p3)

}
