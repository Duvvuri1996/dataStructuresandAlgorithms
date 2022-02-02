/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

type Array struct{
    A []int
}

func mergeSort(arr1 Array, arr2 Array) Array{
    length := len(arr1.A)+len(arr2.A)
    var i, j, k int
    arr3 := Array{make([]int, length)}
    i,j,k = 0,0,0
    for i<len(arr1.A) && j<len(arr2.A) {
        if arr1.A[i] < arr2.A[j] {
            arr3.A[k] = arr1.A[i]
            i++
            k++
        } else {
            arr3.A[k] = arr2.A[j]
            k++
            j++
        }
    }
    for ;i<len(arr1.A); i++{
        arr3.A[k] = arr1.A[i]
        k++
    }
    for ;j<len(arr2.A); j++{
        arr3.A[k] = arr2.A[j]
        k++
    }
    return arr3
}

func main() {
    fmt.Println("Hello World")
    arr1 := Array{[]int{2,5,8,10,15}}
    arr2 := Array{[]int{3,4,7,11,18}}
    var arr3 Array
    
    arr3 = mergeSort(arr1, arr2)
    fmt.Println(arr3.A)
}