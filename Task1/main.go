package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	fmt.Println("Початковий масив", arr)
	mapForClear := make(map[int]int)
	for _, val := range arr {
		mapForClear[val] = val
	}
	var result []int
	for val := range mapForClear {
		result = append(result, val)
	}
	fmt.Println("Очищений масив", result)
}
