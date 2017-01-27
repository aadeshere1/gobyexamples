package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)

	for i, num := range nums {
		if num == 4 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "ball"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for _, k := range kvs {
		fmt.Println("value:", k)
	}

	for i, c := range "मेरो नाम आदेश हो " {
		fmt.Println(i, c)
	}
}
