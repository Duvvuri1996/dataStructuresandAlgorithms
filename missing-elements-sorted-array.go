/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"


func singleElement(arr [10]int){
    var i, sum int = 0,0
    sumOfN := (arr[len(arr)-1]*(arr[len(arr)-1]+1))/2
    for i=0;i<len(arr);i++ {
        sum += arr[i]
    }
    if sum < sumOfN{
        fmt.Println(sumOfN-sum)
    }else{
        fmt.Println("No Missing Element")
    }
}

func singleElement2(arr [10]int){
    var i int = 0
    diff := arr[0]-0
    for ;i<len(arr); i++{
        if arr[i]-i != diff {
            fmt.Println("The missing element is: ", diff+i)
            break;
        }
    }
}

func mulitpleElements(arr [10]int){
    var i int = 0
    diff := arr[0]-0
    for ;i<len(arr); i++{
        if arr[i]-i != diff {
            for diff<arr[i]-i{
                fmt.Println("Element is: ", i+diff)
                diff++
            }
        }
    }
}

func multiElemUnsorted(arr []int){
    var l, h int = findMaxMin(arr)
    var i,j int = 0,l
    hashArr := make([]int, h+1)
    for ;i<len(arr);i++ {
        hashArr[arr[i]]++
    }
    for ;j<=h;j++{
        if hashArr[j] == 0{
            fmt.Println("Missing element is: ", j)
        }
    }
    fmt.Println(hashArr)
}

func findMaxMin(arr []int)(int, int){
    var min, max int = arr[0], arr[0]
    for _, value := range arr{
        if value < min {
            min = value
        } else if value > max{
            max = value
        }
    }
    return min, max
}

func main() {
    arr := []int{16,5,14,15,4,9,10,11,6,7}
    fmt.Println("=============")
    fmt.Println("Missing Element is: ")
    multiElemUnsorted(arr)
}