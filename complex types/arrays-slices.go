
package main

import "fmt"

func main() {

	fmt.Println("Testing complex types")

	// declaring an empty array of strings
	//var days []string // if you uncomment this, code below will fail on line 14
	
	// declaring an array with elements
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	fmt.Println("The first day of week is", days[0])

	for i := 0; i < 7; i++ {
        fmt.Println("Today is", days[i], "and counter is", i)
    }

	weekdays := days[0:5]
	fmt.Println("Weekdays are", weekdays)

}

