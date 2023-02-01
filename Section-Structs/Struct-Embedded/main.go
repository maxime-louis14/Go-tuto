package main

import "fmt"

type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Level int

}

func main() {
	u := User{
		Name:  "Makab",
		Email: "tasoeurellebaslebeures@fdp.fr",
	}
	fmt.Printf("User=%v\n", u)

	a := Admin{
		Level: 2,
		User: User{
			Name: "Nico",
			Email: "nico@tine.fr",
		},
	}
	a.Name = "Nico"
	a.Email = "nico@tine.fr"
	fmt.Printf("Admin=%v\n", a)
	fmt.Printf("Admin=%v, level=%v\n", a.Name, a.Level)
}