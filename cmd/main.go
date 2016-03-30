package main
import (
	"fmt"
	"github.com/dooman87/go-intro"
	"time"
)

func main() {
	fmt.Println("Hello world!\n")

	//Types
	fmt.Println(intro.GetGreeting())
	fmt.Println(intro.GetInferredGreeting())

	//Interfaces
	fmt.Println(greetingFactory().Greet())

	//Closures
	n := 3
	getNumber := func() int {
		return n
	}
	n = 4
	fmt.Printf("Sum: %d", intro.Sum(getNumber, getNumber))

	//Collections
	intro.Slices()
	intro.Maps()

	//Go routines
	go intro.Go(1, 1000)
	go intro.Go(2, 500)
	time.Sleep(10 * time.Second)

	//Channels
	ch := make(chan int)
	go intro.Casino(ch)
	go intro.Gambler(ch)
	time.Sleep(10 * time.Second)

	//Standard library
	intro.RunServer()
}

func greetingFactory() intro.Greeter {
	return new(intro.MachoMan)
}
