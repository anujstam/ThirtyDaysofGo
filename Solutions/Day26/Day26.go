package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var d1 int;
    var m1 int;
    var y1 int;
    var d2 int;
    var m2 int;
    var y2 int;
    var fine int;

    fmt.Scanf("%d", &d1)
    fmt.Scanf("%d", &m1)
    fmt.Scanf("%d", &y1)
    fmt.Scanf("%d", &d2)
    fmt.Scanf("%d", &m2)
    fmt.Scanf("%d", &y2)
    
    if (y1<y2){
        fine = 0
    }else if (y1>y2){
        fine = 10000
    }else if (m1>m2){
        fine = 500 * (m1-m2)
    }else if (d1>d2){
        fine = 15 * (d1-d2)
    }else{
        fine = 0
    }
    

    fmt.Printf("%d\n",fine)
    
}

