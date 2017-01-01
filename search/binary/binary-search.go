package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func binarySearch(list []int, num int) (int, error) {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)
		guess := list[mid]

		if guess == num {
			return mid, nil
		} else if guess > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1, errors.New("Could not find the number in the list.")
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

	// Get the number to find index of.
	fmt.Print("Enter the number to find it's index in the list: ")
	input, err = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	index, err := binarySearch(list, number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(index)
}
