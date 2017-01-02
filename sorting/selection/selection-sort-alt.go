package main

import "fmt"

func selectionSortAlt(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i

		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}

		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
}

func main() {
	nums := []int{1, -1, 23, -2, 23, 123, 12, 1}
	selectionSortAlt(nums)
	fmt.Println(nums)
}
