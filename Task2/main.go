package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string
	var min int32
	var max int32
	dividedString := strings.Fields(input) // Розділення суцільної стрінги на кусочки
	for i, val := range dividedString {
		particle, _ := strconv.ParseInt(val, 10, 32) // Перетворення стрінги в int32
		if i == 0 {
			max = int32(particle)
			min = int32(particle)
		}
		if int32(particle) < min {
			min = int32(particle) //Знаходження мінімуму
		}
		if int32(particle) > max {
			max = int32(particle) //Знаходження максимуму
		}
	}
	minString := strconv.FormatInt(int64(min), 10) // Перетворення int32 в стрінг
	maxString := strconv.FormatInt(int64(max), 10)
	if max == min { // Умова для ситуації з одним символом
		result = (maxString)
	} else {
		result = (maxString + " " + minString)
	}
	fmt.Println(result)
}
