package main

import "fmt"

const p string = "death & taxes"

const (
	PI       = 3.14
	Language = "Go"
)

const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

func main() {
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

// a CONSTANT is a simple unchanging value
