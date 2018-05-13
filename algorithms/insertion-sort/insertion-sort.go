package main

import "fmt"

//InsertionSort algorithm based on Cormen's book
func InsertionSort(s []int) {
	for j := 1; j < len(s); j++ {
		key := s[j]
		i := j - 1
		for i > 0 && s[i] > key {
			s[i+1] = s[i]
			i = i - 1
		}
		s[i+1] = key
	}
	fmt.Printf("sorted: s is %v\n", s)

}

func main() {
	s := []int{1, 2, 66, 66, 3, 5, 999, 324}
	fmt.Printf("Befor sorting: s is %v\n", s)
	InsertionSort(s)
}
