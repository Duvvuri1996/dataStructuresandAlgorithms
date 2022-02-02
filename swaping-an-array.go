/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

func swapingArray(arr []int){
    var i int
    i = 0
    for j:=len(arr)-1; i<j; j--{
        temp := arr[i]
        arr[i] = arr[j]
        arr[j] = temp
        i++
    }
}

func main() {
    fmt.Println("Hello World")
    arr := []int{5,2,7,4,9,8}
    swapingArray(arr)
    fmt.Println(arr)
}