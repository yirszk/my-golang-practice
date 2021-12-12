package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	functionClosure()
}

func functionClosure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*1))
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func mutatingMap() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func Map() {
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex
	fmt.Println(m)

	m = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39667,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func Range() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)
	fmt.Println(pow2)
	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}
	fmt.Println(pow2)
	for i, value := range pow2 {
		fmt.Printf("%d %d\n", i, value)
	}
}

func appendToSlice() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func sliceOfSlice() {
	board := [][]string {
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board)

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func createSlice() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func slice() {
	primes := [7]int{2, 3, 5, 7, 11, 13, 15}
	var s []int = primes[1:5]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	s = []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}
	fmt.Println(s1)

	var sl []int
	sl = s[1:4]
	fmt.Println(sl)

	sl = s[:2]
	fmt.Println(sl)

	sl = s[1:]
	fmt.Println(sl)

	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func pointer() {
	i, j := 42, 2701

	p := &i
	fmt.Println(p, *p, i, j) // 0x140000a2008 42 42 2701

	*p = 21
	fmt.Println(p, *p, i, j) // 0x140000a2008 21 21 2701

	p = &j
	*p = *p / 37
	fmt.Println(j) // 73
}

type Vertex struct {
	X, Y int
}

func Struct() {
	v := Vertex{1,2}
	v.X = 4
	fmt.Println(v)
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func structLiteral() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p  = &Vertex{1, 2}
	)
	fmt.Println(v1, p, v2, v3)
}

func Array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}