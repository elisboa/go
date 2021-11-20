package main

import "fmt"

func main() {

	fmt.Println("Integer overflow using a for loop")

    // integer overflow test
    var myint int8
    for i := 0; i < 129; i++ {
        myint += 1
    }
    fmt.Println(myint) // prints out -127

}
