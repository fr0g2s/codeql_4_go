package main

import "fmt"

type MyStruct struct {
    a int
}

func (s MyStruct) valueMethod() {
    s.a = 5
    fmt.Println("=== valueMethod ===")
    fmt.Printf("a = %d\n", s.a)
}

func main() {
    my_s := MyStruct{}

    my_s.a = 10

    my_s.valueMethod()
    
    fmt.Println("=== main ===")
    fmt.Printf("a = %d\n", my_s.a)
}
