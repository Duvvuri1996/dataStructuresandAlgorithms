package main

import "fmt"

type Array struct {
	a      []int
	length int
	size   int
}

func main() {
	arr := []int{12, 13, 14, 15, 16, 17, 18, 19}
	res := rbinarySearch(arr, 0, len(arr)-1, 18)
	fmt.Println("Hello World", res)
}

func rbinarySearch(arr []int, l int, h int, key int) int {
	fmt.Println(l, h)
	if l <= h {
		mid := int((l + h) / 2)
		if arr[mid] == key {
			return mid
		} else if arr[mid] < key {
			return rbinarySearch(arr, mid+1, h, key)
		} else if arr[mid] > key {
			return rbinarySearch(arr, l, mid-1, key)
		}
	}
	return -1
}
