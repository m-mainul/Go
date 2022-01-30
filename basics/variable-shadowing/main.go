package main

import (
	"fmt"
)

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7) // max is now the result, not the function
	fmt.Println(max)
}

// don't do this; bad coding practice to shadow variables
