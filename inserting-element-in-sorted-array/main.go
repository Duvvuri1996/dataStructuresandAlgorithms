package main

import (
	"fmt"
	"reflect"
)

func insertSort(arr *[]int, element int) {
	fmt.Println(*arr)
	i := len(*arr) - 1
	*arr = append(*arr, 0)
	for (*arr)[i] > element {

		(*arr)[i+1] = (*arr)[i]
		i--
	}
	(*arr)[i+1] = element
}

func main() {
	a := []int{3, 4, 5, 6, 8, 9}
	arr := make([]int, len(a), 10)
	fmt.Println(reflect.ValueOf(arr).Kind())
	for index := range a {
		arr[index] = a[index]
	}
	insertSort(&arr, 7)
	fmt.Println(arr)
}
