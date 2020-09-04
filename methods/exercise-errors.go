package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}


	z := 1.0
	const Eps = 1e-9

	cnt := 0

	for ; math.Abs(z * z - x) > Eps ; cnt++ {
		z -= (z * z - x) / (2 * z)
	}

	fmt.Printf("cnt: %v\n", cnt)

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
