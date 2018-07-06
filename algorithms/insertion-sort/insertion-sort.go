package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//InsertionSort algorithm based on Cormen's book
func InsertionSort(s []int) []int {
	innerS := make([]int, len(s))
	copy(innerS, s)
	for j := 1; j < len(innerS); j++ {
		key := innerS[j]
		i := j - 1
		for i >= 0 && s[i] > key {
			innerS[i+1] = innerS[i]
			i = i - 1
		}
		innerS[i+1] = key
	}
	return innerS

}

func main() {

	slice := generateSlice(10)
	//fmt.Println("\n--- Unsorted --- \n\n", slice)
	//fmt.Println("\n--- Sorted ---\n\n", InsertionSort(slice), "\n")
	defer timeTrack(time.Now(), "InsertionSort")
	sortedSlice := InsertionSort(slice)
	fmt.Print(slice)
	fmt.Print(sortedSlice)

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
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
