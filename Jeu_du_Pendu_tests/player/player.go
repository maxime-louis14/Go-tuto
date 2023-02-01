package player

type Avatar struct {
	Url string
}

type Gamer struct {
	Name     string
	Age      int
	Avatar   Avatar
	password string
}
