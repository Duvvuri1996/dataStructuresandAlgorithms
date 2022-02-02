/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

func linearSearch(arr []int, key int)int{
    for index, element := range arr{
        if element == key{
            temp := arr[index];
            arr[index] = arr[index-1];
            arr[index-1] = temp;
            return index
        }
    }
    return -1;
}

func BlinearSearch(arr []int, key int) int {
    for index, element := range arr{
        if element == key {
             temp := arr[index]
             arr[index] = arr[0]
             arr[0] = temp
             return index
        }
    }
    return -1;
}

func main() {
    arr := []int{2,3,1,4,6,7}
    //Bres := BlinearSearch(arr, 1)
    res := linearSearch(arr, 1)
    fmt.Println("Hello World", res)
    fmt.Println(arr)
}