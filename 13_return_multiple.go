package main

import "fmt"

func main() {
	a, b, d := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)
	_, _, c := vals()
	fmt.Println(c)
}

func vals() (int, int, int) {
	return 3, 7, 5
}
