package main
import "fmt"

func bsort(arr []int) ([]int,int){
    noswaps := 0
    swapcount := 0
    var temp int
    for i:=0;i<len(arr);i++{
        noswaps = 0
        for j:=0;j<len(arr)-i-1;j++{
            if arr[j]>arr[j+1]{
                temp = arr[j]
                arr[j] = arr[j+1]
                arr[j+1] = temp
                noswaps = 1
                swapcount += 1
            }
        }
        if noswaps == 0{
            break
        }
    }
    return arr,swapcount
}

func main() {
    var size int
    var swaps int
    fmt.Scanf("%d", &size)
    var arr = make([]int,size)
    for i:=0;i<size;i++{
        fmt.Scanf("%d", &arr[i])
    }

    arr,swaps = bsort(arr)

    fmt.Printf("Array is sorted in %d swaps.\n",swaps)
    fmt.Printf("First Element: %d\n",arr[0])
    fmt.Printf("Last Element: %d\n",arr[size-1])
}

