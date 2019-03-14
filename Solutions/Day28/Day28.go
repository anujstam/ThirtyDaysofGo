package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "regexp"
    "sort"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    N := int32(NTemp)
    M := make(map[string]string)
    var Names []string

    for NItr := 0; NItr < int(N); NItr++ {
        firstNameEmailID := strings.Split(readLine(reader), " ")

        firstName := firstNameEmailID[0]

        emailID := firstNameEmailID[1]

        M[emailID] = firstName
    }
    for key, value := range M {
        matched, _ := regexp.MatchString("@gmail.com", key)
        if matched{
            Names = append(Names, value)
        }
    }
    sort.Strings(Names)

    for _,name := range Names{
        fmt.Printf("%s\n",name)
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
