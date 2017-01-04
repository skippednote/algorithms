package main

import "fmt"

func quickSort(list []int) []int {
	if len(list) == 0 {
		return []int{}
	} else if len(list) == 1 {
		return list
	}

	first := list[0]
	rest := list[1:]

	var left, right []int
	for _, v := range rest {
		if v <= first {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left, right = quickSort(left), quickSort(right)
	var new []int
	new = append(new, left...)
	new = append(new, first)
	new = append(new, right...)

	return new
}

func main() {
	list := []int{5, 7, 2, 2, 8, 3, 9}
	sorted := quickSort(list)
	fmt.Println(sorted)
}
