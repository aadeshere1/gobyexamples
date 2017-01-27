package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}
func main() {
	res := plus(1, 3)
	fmt.Println("1+3=", res)

	res = plusPlus(res, 5, 6)
	fmt.Println("4 + 5 + 6 = ", res)

}
