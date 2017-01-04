package main

import "fmt"

func sum(numbers []int, total *int) int {
	if len(numbers) <= 0 {
		return *total
	}
	*total += numbers[0]

	ref := numbers[1:]
	return sum(ref, total)

}

func length(numbers []int, count *int) int {
	if len(numbers) <= 0 {
		return *count
	}
	*count++

	ref := numbers[1:]
	return length(ref, count)
}

func max(numbers []int, high *int) int {
	if len(numbers) <= 0 {
		return *high
	}
	ref := numbers[1:]
	if numbers[0] > *high {
		*high = numbers[0]
		return max(ref, high)
	}
	return max(ref, high)
}

func main() {
	total := 0
	result := sum([]int{1, 2, 3, 4, 5}, &total)
	fmt.Println(result)

	count := 0
	result = length([]int{1, 2, 3, 4, 5}, &count)
	fmt.Println(result)

	high := 0
	result = max([]int{1, 2, 3, 4, 5}, &high)
	fmt.Println(result)
}
