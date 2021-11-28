package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return x, y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2

var (
	ToBe 	bool 		= false
	MaxInt 	uint 		= 1<<64 - 1
	z		complex128	= cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

const (
	Big = 1 << 100
	Small = Big >> 99
	)

func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64 {
	return x * 0.1
}


func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(12,34))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(9876))
	c, python, java := true , false,  "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var x,  y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
	v := 42
	fmt.Printf("v is of type %T\n", v)

	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(Small)
}