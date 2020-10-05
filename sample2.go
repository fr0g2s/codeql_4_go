package main

import "fmt"

type Person struct{
    age int
}

func (p People) valueMethod() {
    p.age = 5
    fmt.Println("=== valueMethod ===")
    fmt.Printf("age = %d\n", p.age)
}

func main() {
    p := Person{}

    p.age = 200

    p.valueMethod()

    fmt.Println("=== main ===")
    fmt.Printf("age = %d\n", p.age)
}
