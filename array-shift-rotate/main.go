package main

import "fmt"

func leftRotate(arr []int) {
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp
}

func rightRotate(arr []int) {
	temp := arr[len(arr)-1]
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = temp
}

func main() {
	fmt.Println("Hello World")
	arr := []int{3, 4, 5, 7, 6}
	//leftRotate(arr)
	rightRotate(arr)
	fmt.Println(arr)
}
