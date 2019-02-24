package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func mults(n int32){
    var i int32
    for i=1;i<11;i++{
        fmt.Println(n,"x",i,"=",n*i)
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    var n int32
    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n = int32(nTemp)
    mults(n)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}