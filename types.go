package intro

var greeting string

func init() {
	greeting = "Hello Go!\n---------\n"
}

func GetGreeting() string {
	return greeting
}

func GetInferredGreeting() string {
	newGreeting := "Well Well Well\n"
	return newGreeting
}
