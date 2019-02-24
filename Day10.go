package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func maxones(n int32) int{
    onecount := 0
    temp := n
    curr := 0
    for temp>0{
        if temp%2==1{
            curr+=1
        }else{
            if curr>onecount{
                onecount = curr
            }
            curr = 0
        }
        temp/=2
    }
    if curr>onecount{
        onecount = curr
    }
    return onecount
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)
    fmt.Print(maxones(n))
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
