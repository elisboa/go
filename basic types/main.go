// the first statement of every go source file
// must be a package declaration. If we aren't doing anything
// fancy, this tends to be package main.
package main

// We then want to use the fmt package
// which features a `print` function - Println
import "fmt"

// We then need to define our main function.
// Think of this as the entry point to our Go
// program
func main() {
    // within this main function, we then
    // want to call a function within the fmt
    // package called Println() in order to print
    // out `Hello World`
    fmt.Println("Hello World")

    // integer overflow test
    var myint int8
    for i := 0; i < 129; i++ {
        myint += 1
    }
    fmt.Println(myint) // prints out -127


    // Simple bool verification
    var amazing bool
    amazing = true
    if amazing {
      fmt.Println("Subscribe to our channel!")
    }

    // Testing 'If' statement 
    var isTrue bool = true
    var isFalse bool = true
    // AND
    if isTrue && isFalse {
      fmt.Println("Both Conditions need to be True")
    }
    // OR - not exclusive
    if isTrue || isFalse {
      fmt.Println("Only one condition needs to be True")
    }

    var myName string
    myName = "Eduardo Lisboa"
    fmt.Println("My name is", myName)

    const meaningOfLife = 42
    fmt.Println("The meaning of life is", meaningOfLife)

}

