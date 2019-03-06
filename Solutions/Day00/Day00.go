package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    inp,_ := reader.ReadString('\n')
    fmt.Println("Hello, World.")
    fmt.Println(inp)
}
