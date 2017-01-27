package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)
	// `const` can appear anywhere a var statement can
	const n = 500000

	const d = 3e20 / n

	fmt.Println(d)
	// numeric constant has no type until its given.
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
