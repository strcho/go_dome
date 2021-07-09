package main



import (
	"fmt"
	"math"
)

func add(x int, y int) (int, string) {
	return x + y, "ok!"
}

func main() {
	a, b := add(44, 99)
	fmt.Printf("?? %d, %s\n", a, b)
	fmt.Printf("Now you have %f problems.\n", math.Sqrt(7))
}
