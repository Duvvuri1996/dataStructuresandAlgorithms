package main

import "fmt"

func reversing1(arrA []int, length int) {
	arrB := make([]int, length)
	var i, j, k int
	j = 0
	for i = length - 1; i >= 0; i-- {
		fmt.Println(i, j)
		arrB[j] = arrA[i]
		j++
	}
	for k = 0; k < length; k++ {
		arrA[k] = arrB[k]
	}
}

func main() {
	fmt.Println("Hello World")
	arrA := []int{3, 2, 89, 45, 4}
	length := len(arrA)

	reversing1(arrA, length)

	fmt.Println(arrA)
}
