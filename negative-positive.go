/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

func checkSorted(arr []int) bool {
    for i:=0; i<len(arr)-1; i++ {
        if arr[i] > arr[i+1]{
            return false
        }
    }
    return true
}

func negativePositive(arr []int){
    var i, j int
    i,j = 0,len(arr)-1
    for i<j{
        for arr[i]<0 {
        fmt.Println(i, arr[i])
        i++;
    }
    for arr[j]>0{
        fmt.Println(j, arr[j])
        j--
    }
    if(i<j){
        temp := arr[i]
        arr[i] = arr[j]
        arr[j] = temp
    }
    }
}

func main() {
    fmt.Println("Hello World")
    arr := []int{4,-5,-6,-7,8}
    negativePositive(arr)
    fmt.Println(arr)
}