package main

import "fmt"

func main() {

        //set slice values
        arr := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

        //Call func to find min value and output `min` values in stdout
        fmt.Println("Minimal value is:", Min(arr))
}
func Min(arr []int) int {

        //set start value to 'min' variable
        min := arr[0]

        //Use `for-range` loop to compare values each to other
        for _, value := range arr {
                if value < min {      // if current value less that current `min` variable
                        min = value       // then replace `min` on current value
                }
        }

        return min
}
