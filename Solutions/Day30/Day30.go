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

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nk := strings.Split(readLine(reader), " ")

        nTemp, err := strconv.ParseInt(nk[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        kTemp, err := strconv.ParseInt(nk[1], 10, 64)
        checkError(err)
        k := int32(kTemp)

        /*
        Logic is as follows:
        For any odd K, the binary repersentation will (Somebinary)1
        The number K-1 will then be (Somebinary)0, and hence K & (K-1) will return (Somebinary)0 = K-1

        For even K, since (K-1) = (K-1) &  (K | (K-1)), we need to check the value of (K | (K-1))
        If ( K | (K-1)) is less than or equal to N, then it is in S, and so we can use the value so the result is still K-1
        
        If it is more than N then it means that K-1 can't be obtained, so the answer is the same as having K-1 as the upper limit
        and since K-1 is odd, K-2 is even and that will be the answer 

        */

        if (k%2 == 1){
            fmt.Printf("%d\n",(k-1))
        }else{
            if((k-1)|k <= n){
                fmt.Printf("%d\n",(k-1))
            }else{
                fmt.Printf("%d\n",(k-2))
            }
        }
    }
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
