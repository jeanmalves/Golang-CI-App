package main

import (
    "fmt"
    "github.com/jeanmalves/operation/addition"
)

func main() {
    var num1 = 5
    var num2 = 5
    fmt.Printf("A soma de %d + %d é %d\n", num1, num2, addition.Add(num1, num2))
}

