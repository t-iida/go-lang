package main

import "fmt"

func main() {

    var intArray = []int {9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

    for i := 0; i < len(intArray); i++ {
        fmt.Print(intArray[i]);
    }
    fmt.Print("\n")

    for i := 0; i < len(intArray); i++ {
        intArray[i] = i
    }

    for i := 0; i < len(intArray); i++ {
        fmt.Print(intArray[i]);
    }
    fmt.Print("\n")

}
