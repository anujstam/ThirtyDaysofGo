package main
import (
    "fmt"
    "math"
)

func isprime(n int) bool{
    lim := int(math.Pow(float64(n),0.5))
    if (n==1){
        return false
    }
    for i:=2;i<=lim;i++{
        if(n%i==0){
            return false
        }
    }
    return true
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var t int;
    var n int;
    fmt.Scanf("%d\n",&t)
    for i:=0;i<t;i++{
        fmt.Scanf("%d",&n)
        if (isprime(n)){
            fmt.Printf("Prime\n")
        }else{
            fmt.Printf("Not prime\n")
        }
    }
}

