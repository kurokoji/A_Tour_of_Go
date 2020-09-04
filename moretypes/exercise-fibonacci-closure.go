package main

import "fmt"

func fibonacci() func() int {
	sum1 := 0
	sum2 := 1

	return func() int {
		sum := sum1 + sum2
		sum1 = sum2
		sum2 = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
