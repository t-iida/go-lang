/*
 * Pointer sample
 *
 * Passing parameter by using pointer.
 *
 */
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

}
