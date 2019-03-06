package main
import(
   "bufio" 
   "fmt"
   "os"
)

func stsplit(s string){
    length := len(s)
    for i:=0;i<length;i++{
        fmt.Print(string(s[i]))
    }
    for j:=1;j<length;j+=2{
        fmt.Print(string(s[j]))
    }
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var looper int
    reader := bufio.NewReader(os.Stdin)
    fmt.Scanln(&looper)
    for i:=0;i<looper;i++{
        txt,_ := reader.ReadString('\n')
        length := len(txt)
        if txt[length-1]=='\n'{
            length-=1
        }
        for i:=0;i<length;i+=2{
            fmt.Print(string(txt[i]))
        }
        fmt.Print(" ")
        for j:=1;j<length;j+=2{
            fmt.Print(string(txt[j]))
        }
        fmt.Print("\n")
    }

}

