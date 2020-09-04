package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	const Eps = 1e-9

	cnt := 0

	for ; math.Abs(z * z - x) > Eps ; cnt++ {
		z -= (z * z - x) / (2 * z)
	}

	fmt.Printf("cnt: %v\n", cnt)

	return z
}

func main() {
	fmt.Println(Sqrt(144))
}
