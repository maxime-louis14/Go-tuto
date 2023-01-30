package main

import (
	"fmt"

	"training.go/Jeu_du_Pendu_tests/player"
)

func main() {
	var p1 player.Player
	p1.Name = "Maka"
	p1.Age = "24"

	fmt.Printf("Player 1 %v\n", p1)

}
