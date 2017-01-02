package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findSmallest(list []int) int {
	smallest := list[0]
	smallestIndex := 0

	for i := 1; i < len(list); i++ {
		if list[i] < smallest {
			smallest = list[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

func selectionSort(list, sortedList *[]int, loop int) []int {
	for i := 0; i < loop; i++ {
		smallest := findSmallest(*list)

		var listRef = *list
		*sortedList = append(*sortedList, listRef[smallest])
		*list = append(listRef[:smallest], listRef[smallest+1:]...)
	}

	return *sortedList
}

func main() {
	// Get the list of numbers.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a list of numbers separated by comma: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	var list []int
	for _, v := range strings.Split(input, ",") {
		v, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, v)
	}

	sortedList := []int{}
	selectionSort(&list, &sortedList, len(list))
}
