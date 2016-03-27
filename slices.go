package intro
import "fmt"

func Slices() {
	slice := make([]int, 0)
	slice = append(slice, 9, 8, 7, 6, 5, 4, 3, 2, 1)

	fmt.Printf("%v\n", slice[8])
	fmt.Printf("%v\n", slice[3:5])
	fmt.Printf("%v\n", slice[7:])

	//Delete 2nd element
	slice = append(slice[:1], slice[2:]...)
	fmt.Printf("%v\n", slice)

	fmt.Printf("len = %d, cap = %d\n", len(slice), cap(slice))

	//fmt.Printf("%v\n", slice[100])
}
