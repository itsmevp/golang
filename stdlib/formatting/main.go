package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  // Reading from standard input
  reader := bufio.NewReader(os.Stdin)
  s, _ := reader.ReadString('\n')
  fmt.Println(s)

  // Formatting and precision
  x := 10.0
  fmt.Printf("%.1f\n",x)
  fmt.Printf("%f\n", x)
  fmt.Printf("%11.2f\n", x)
}