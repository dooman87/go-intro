package intro

type Greeter interface {
	Greet() string
}

type MachoMan struct {
	Age int
}

func (mr *MachoMan) Greet() string {
	return "Aloha!"
}
