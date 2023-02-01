package main

import (
	"fmt"

	// Tu importes de dossier player par le go.mod a la racine du projet.
	// avec go mod nom que tu veux lui donner
	"github.com/maxime-louis14/Go-tuto/Section-Structs/Struct-declaration/player"
)

func main() {
	// Ici tu crées un variable pour appeler le struct
	//dans l'autre dossier

	var p1 player.Player
	p1.Name = "Maka"
	p1.Age = 24
	// A la fin du fmt.Printf (p1), tu le déclares
	fmt.Printf("Player 1 %v\n", p1 )
	fmt.Printf("p1.Name=%v, p1.Age=%v\n", p1.Name, p1.Age)

	// Avatar
	a := player.Avatar{"http://avatar.jpg"}
	fmt.Printf("avatar : %v\n", a)

	p3 := player.Player{
		Name: "MAka",
		Age:  25,
		Avatar: player.Avatar{
			Url: "http://url.com",
		},
	}
	fmt.Printf("Player 3 : %v\n", p3)

	p4 := player.New("Hatmos")
	fmt.Printf("Player 4 : %v\n", p4)
}
