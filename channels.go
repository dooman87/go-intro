package intro
import (
	"fmt"
	"math/rand"
	"time"
)

func Casino(ch chan int) {
	var casino, gambler int

	for {
		casino = rand.Intn(6)

		gambler = <- ch
		if casino > gambler {
			fmt.Printf("Casino won! (%d:%d)\n", casino, gambler)
		} else if casino < gambler {
			fmt.Printf("Gambler won! (%d:%d)\n", casino, gambler)
		} else {
			fmt.Printf("Draw (%d:%d)\n", casino, gambler)
		}
	}
}

func Gambler(ch chan int) {
	var myScore int

	for {
		myScore = rand.Intn(6)
		ch <- myScore
		time.Sleep(1 * time.Second)
	}
}