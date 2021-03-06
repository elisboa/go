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

    var myName string
    var myYear uint16
    
    myName = "Eduardo Lisboa"
    myYear = 1982
    
    fmt.Println("My name is", myName, "and I was born in", myYear)

    const meaningOfLife = 42
    fmt.Println("The meaning of life is", meaningOfLife)

}
