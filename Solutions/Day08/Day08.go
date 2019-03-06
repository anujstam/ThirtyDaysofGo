package main
import (
    "fmt"
    "bufio"
    "os"
    "io"
)

func givenumber(input io.Reader) error {
    var phoneNum int
    fmt.Scan(&phoneNum)

    phoneBook := make(map[string]string)
    for i := 0; i < phoneNum; i++ {
        var name, phone string
        fmt.Scan(&name, &phone)
        phoneBook[name] = phone
    }

    s := bufio.NewScanner(input)
    for s.Scan() {
        nameToPrint := s.Text()
        if p, ok := phoneBook[nameToPrint]; ok {
            fmt.Printf("%s=%s\n", nameToPrint, p)
        } else {
            fmt.Println("Not found")
        }
    }

    return nil
}

func main() {
    givenumber(os.Stdin)
}

