package main

import (
	"fmt"
	"math/rand"
	"time"
)

//InsertionSort algorithm based on Cormen's book
func InsertionSort(s []int) []int {
	for j := 1; j < len(s); j++ {
		key := s[j]
		i := j - 1
		for i >= 0 && s[i] > key {
			s[i+1] = s[i]
			i = i - 1
		}
		s[i+1] = key
	}
	return s

}

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	fmt.Println("\n--- Sorted ---\n\n", InsertionSort(slice), "\n")
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
