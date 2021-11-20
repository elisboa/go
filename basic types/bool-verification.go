package main

import "fmt"

func main() {

	fmt.Println("Boolean verification")

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

}

