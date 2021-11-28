package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func forStatement(sum int) int {
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x,n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func switchStatement () {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switchEvaluationOrder () {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func deferStatement () {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func stackingDefer () {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	fmt.Println(forStatement(10))
	fmt.Println(sqrt(2))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
	switchStatement()
	switchEvaluationOrder()
	deferStatement()
	stackingDefer()
}