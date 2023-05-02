package main

import "fmt"

func hitungPoin(poin ...int) int {
	totalPoin := 0
	for _, nilai := range poin {
		totalPoin += nilai
	}
	return totalPoin
}

func main() {
	totalPoin := hitungPoin(16,17,18,19,20)
	fmt.Println(totalPoin)
}