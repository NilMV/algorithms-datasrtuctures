package main

import (
	"log"
	"math/rand"
	"time"
)

//BubbleSort algorithm based on Cormen's book
func BubbleSort(s []int) []int {
	swtch := true
	for i := 1; i < len(s)-1 && swtch == true; i++ {
		swtch = false
		for j := 0; j < len(s)-i; j++ {
			if s[j] > s[j+1] {
				s[j+1], s[j] = s[j], s[j+1]
				swtch = true
			}
		}
	}
	return s

}

func main() {

	slice := generateSlice(100000)
	//fmt.Println("\n--- Unsorted --- \n\n", slice)
	//defer timeTrack(time.Now(), "BubbleSort")
	//fmt.Println("\n--- Sorted ---\n\n", BubbleSort(slice), "\n")
	defer timeTrack(time.Now(), "BubbleSort")
	BubbleSort(slice)
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
