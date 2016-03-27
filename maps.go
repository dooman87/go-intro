package intro
import "fmt"

func Maps() {
	m := make(map[int]string, 0)

	m[1] = "one"
	m[2] = "two"
	m[3] = "three"

	fmt.Printf("%v\n", m)

	for k, v := range m {
		fmt.Printf("Key: %d, Value: %s\n", k, v)
	}

	delete(m, 2)
	fmt.Printf("%v\n", m)
}
