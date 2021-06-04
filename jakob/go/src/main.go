
package main

import (
    "fmt"
)

func f1(s string) {
    s = "f1"
}

func f2(s *string) {
    *s = "f2"
}

func main() {
    var s string = "main"
    fmt.Println("Value bevore function:", s)
    f1(s)
    fmt.Println("Value after function:", s)
    fmt.Println("Value bevore function:", s)
    f2(&s)
    fmt.Println("Value after function:", s)
}
