package main

import "fmt"

// Accepts parameter as pointer
func inc(value *int) (int) {
    *value++
    return *value
}



func main() {

    var intValue int = 0
    fmt.Printf("%v\n", inc(&intValue))
    fmt.Printf("%v\n", inc(&intValue))
    fmt.Printf("%v\n", inc(&intValue))
    fmt.Printf("%v\n", inc(&intValue))
    fmt.Printf("%v\n", inc(&intValue))

    /*
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
    */

}
