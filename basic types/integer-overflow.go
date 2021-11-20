package main

import "fmt"

func main() {

	fmt.Println("8-bit signed integer overflow using a 'for' loop, from 0 to 129")

    // integer overflow test
    var myint int8
    for i := 0; i < 129; i++ {
        myint += 1
    }
    fmt.Println(myint) // prints out -127

}
