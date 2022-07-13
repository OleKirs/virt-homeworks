package main

import "fmt"
import "testing"

func TestMain(t *testing.T) {
	var tr []int
	tr = FilterList()
	if tr[1] != 6 || tr[2] != 9 || tr[9] != 30 {
		s := fmt.Sprintf("Value must be `6, 9, 30`, but here is: %v, %v and %v", tr[1], tr[2], tr[9])
		t.Error(s)
	}
}
