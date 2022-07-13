// Find all integer values to limit value, that may divide on divider without remains
package main

import "fmt"

var set_divider int = 3
var set_limit int = 100

func FilterList() (devide_wo_remains []int) {
	for i := set_divider; i <= set_limit; i += set_divider {
		devide_wo_remains = append(devide_wo_remains, i)
	}
	return
}

func main() {
	//toPrint := FilterList()
	fmt.Printf("Numbers from 1 to `limit` that may divide on `divider` without a remains: \n")
	fmt.Printf("%v", FilterList())
}
