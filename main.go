package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	index := 2
	val := 1000

	res := insert(s, val, index)

	fmt.Println(res)
}

func insert(s []int, v, index int) []int {
	var tempVal int
	s = append(s, tempVal)
	s = append(s[:index+1], s[index:len(s)-1]...)
	s[index] = v
	return s
}