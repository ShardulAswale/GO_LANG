// Merge Sort in Golang
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var x int

	fmt.Println("\nThis an implementation of merge sort")
	fmt.Println("\n------------------------------------------------------\n Choose the order of the elements \n1.Ascending \n2.Descending")
	fmt.Println("\nYour choice is  ")
	fmt.Scan(&x)
	switch x {
	case 1:
		fmt.Println("\nYou choose ascending order")
		slice := generateSlice(20)
		fmt.Print("\nGenertaing a random string please wait")
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print("  *")
		}
		fmt.Println("\n--- Unsorted --- \n\n", slice)
		fmt.Print("\n\n Sorting the array please wait")
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print("  *")
		}
		var start time
		start = time.Now().Unix()
		fmt.Println(start)
		fmt.Println("\n--- Sorted ---\n\n", mergeSorta(slice), "\n")
	case 2:
		fmt.Println("You choose descending order")
		slice := generateSlice(20)
		fmt.Print("\nGenertaing a random string please wait ")
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print("  *")
		}
		fmt.Println("\n--- Unsorted --- \n\n", slice)
		fmt.Print("\n\n Sorting the array please wait")
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print("  *")
		}
		fmt.Println("\n--- Sorted ---\n\n", mergeSortd(slice), "\n")
	}
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func mergeSorta(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return mergea(mergeSorta(left), mergeSorta(right))
}

func mergea(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

func mergeSortd(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merged(mergeSortd(left), mergeSortd(right))
}

func merged(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
