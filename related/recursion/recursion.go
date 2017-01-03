package main

import "fmt"

func sum(numbers []int, counter *int) int {
	if len(numbers) <= 0 {
		return *counter
	}
	*counter += numbers[0]

	ref := numbers[1:]
	return sum(ref, counter)

}

func main() {
	counter := 0
	total := sum([]int{1, 2, 3, 4, 5}, &counter)
	fmt.Println(total)
}
