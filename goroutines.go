package intro
import (
	"fmt"
	"time"
)

func Go(num int, howFast time.Duration) {
	for i:=0; i < 5; i++ {
		fmt.Printf("Number %d\n", num)
		time.Sleep(howFast * time.Millisecond)
	}

	fmt.Printf("Number %d FINISH!\n", num)
}
