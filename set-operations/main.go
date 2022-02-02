package main

import "fmt"

type Array struct {
	A []int
}

//A Union B
func Union(arr1 Array, arr2 Array) Array {
	var i, j, k int = 0, 0, 0
	length := len(arr1.A) + len(arr2.A)
	arr3 := Array{make([]int, length)}
	for i < len(arr1.A) && j < len(arr2.A) {
		if arr1.A[i] < arr2.A[j] {
			arr3.A[k] = arr1.A[i]
			k++
			i++
		} else if arr2.A[j] < arr1.A[i] {
			arr3.A[k] = arr2.A[j]
			k++
			j++
		} else if arr1.A[i] == arr2.A[j] {
			arr3.A[k] = arr1.A[i]
			i++
			j++
			k++
		}
	}
	for ; i < len(arr1.A); i++ {
		arr3.A[k] = arr1.A[i]
		k++
	}
	for ; j < len(arr2.A); j++ {
		arr3.A[k] = arr2.A[j]
		k++
	}
	return arr3
}

//A Intersection B
func Intersection(arr1 Array, arr2 Array) Array {
	var i, j, k int = 0, 0, 0
	length := len(arr1.A) + len(arr2.A)
	arr3 := Array{make([]int, length)}
	for i < len(arr1.A) && j < len(arr2.A) {
		if arr1.A[i] < arr2.A[j] {
			i++
		} else if arr2.A[j] < arr1.A[i] {
			j++
		} else if arr1.A[i] == arr2.A[j] {
			arr3.A[k] = arr1.A[i]
			i++
			j++
			k++
		}
	}
	return arr3
}

//Difference A-B
func Difference(arr1 Array, arr2 Array) Array {
	var i, j, k int = 0, 0, 0
	length := len(arr1.A) + len(arr2.A)
	arr3 := Array{make([]int, length)}
	for i < len(arr1.A) && j < len(arr2.A) {
		if arr1.A[i] < arr2.A[j] {
			arr3.A[k] = arr1.A[i]
			i++
			k++
		} else if arr2.A[j] < arr1.A[i] {
			j++
		} else {
			i++
			j++
		}
	}
	return arr3
}

func main() {
	fmt.Println("Hello World")
	arr1 := Array{[]int{3, 4, 5, 6, 10}}
	arr2 := Array{[]int{2, 4, 5, 7, 12}}

	var arr3 Array
	arr3 = Union(arr1, arr2)
	fmt.Println(arr3.A)
}
