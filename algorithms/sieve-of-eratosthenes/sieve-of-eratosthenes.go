package main

import "fmt"

func printPrimes(n int) {
	prime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		prime[i] = true
	}
	for divisor := 2; divisor*divisor <= n; divisor++ {
		if prime[divisor] {
			for i := 2 * divisor; i <= n; i = i + divisor {
				prime[i] = false
			}
		}
	}
	for i := 2; i <= n; i++ {
		if prime[i] {
			fmt.Println(i)
		}
	}
}

func main() {
	printPrimes(11)
}
