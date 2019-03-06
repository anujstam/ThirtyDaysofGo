package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    // Declare second integer, double, and String variables.
    var ip1 uint64
    var ip2 float64
    var ip3 string


    // Read and save an integer, double, and String to your variables.
    
    fmt.Scan(&ip1)
    fmt.Scan(&ip2)
    _ = scanner
    reader := bufio.NewReader(os.Stdin)
    ip3,_ = reader.ReadString('\n')
    
    // Print the sum of both integer variables on a new line.
    fmt.Println(i+ip1)
    
    // Print the sum of the double variables on a new line.
    fmt.Printf("%.1f\n",d+ip2)
    // Concatenate and print the String variables on a new line
    fmt.Println(s+ip3)
    // The 's' variable above should be printed first.

}
