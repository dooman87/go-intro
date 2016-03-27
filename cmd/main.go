package main
import (
	"fmt"
	"github.com/dooman87/go-intro"
	"os"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello world!\n")
	fmt.Println(intro.GetGreeting())
	fmt.Println(greetingFactory().Greet())

	intro.Slices()

	intro.Maps()

	go intro.Go(1, 1000)
	go intro.Go(2, 500)
	time.Sleep(10 * time.Second)

	ch := make(chan int)
	go intro.Casino(ch)
	go intro.Gambler(ch)
	time.Sleep(10 * time.Second)

	go intro.RunServer()
	pid := os.Getpid();
	fmt.Printf("PID: %d", pid)

	proc, err := os.FindProcess(pid)
	if err != nil {
		log.Fatal(err)
	}
	proc.Wait()
}

func greetingFactory() intro.Greeter {
	return new(intro.MachoMan)
}
