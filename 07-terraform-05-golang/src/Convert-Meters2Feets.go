// Convert size from SI `meters` to Imperial `feets`.
// Run programm, then input size in meters and that convert size to feets.
package main

import (
        "fmt"
        "os"
)

func convert_meters_to_feets(size_in_meters float32) (size_in_feets float32) {

        const meters_to_feet float32 = 0.3048 // describe how mach meters is in 1 feet

        size_in_feets = size_in_meters / meters_to_feet

        return size_in_feets
}

func main() {

        fmt.Println("Input sise (in meters):")

        var input_from_stdin float32
        fmt.Fscan(os.Stdin, &input_from_stdin)

//      Uncomment to debug
//      fmt.Println("Input is: %\n", input_from_stdin)

        fmt.Println("This size is: ", convert_meters_to_feets(input_from_stdin), "feet(s)")
}
