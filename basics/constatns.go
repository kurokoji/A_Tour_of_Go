package main

import "fmt"

// 定数は:=を使っては宣言できない
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
