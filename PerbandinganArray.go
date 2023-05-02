package main

import "fmt"

func main() {
	var array1 = [3]string{
		"a", "c", "d",
	}

	var array2 = [3]string{
		"a", "d", "c",
	}

	temp := make([]int, 0)

	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			temp = append(temp, i)
		}
	}

	fmt.Println("array 1 :", array1)
	fmt.Println("array 2 :", array2)

	for i := 0; i < len(temp); i++ {
		fmt.Println("index ke ", temp[i],"berbeda")
	}
}