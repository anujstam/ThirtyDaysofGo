package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(readLine(reader), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(6) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }
    hourglass(arr)
}

func hourglass(arr [][]int32){
    var max int32;
    var s int32; 
    for i:=1;i<5;i++{
        for j:=1;j<5;j++{
            s = arr[i][j] +
                arr[i-1][j-1]+arr[i-1][j]+arr[i-1][j+1]+
                arr[i+1][j-1]+arr[i+1][j]+arr[i+1][j+1]
            if (i==1 && j==1){
                max = s
            }
            if s>max{
                max=s
            }
        }
    }
    fmt.Print(max)
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

