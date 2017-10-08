package main

import "fmt"

func sum(a int, b int) (c int) {
    c = a + b
    return c
}

func main() {
    c := sum(1, 2)
    fmt.Println(c)
}