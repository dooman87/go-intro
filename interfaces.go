package intro

type Greeter interface {
	Greet() string
}

type MachoMan struct {}

func (mr *MachoMan) Greet() string {
	return "Aloha!"
}
