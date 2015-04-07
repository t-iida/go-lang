/*
 * Array Sample
 */
package main

import "fmt"

func main() {
    var a1 []int = []int{0, 1, 2}
    for i := 0; i < len(a1); i++ {
        fmt.Printf("%v\n", a1[i])
    }
}
