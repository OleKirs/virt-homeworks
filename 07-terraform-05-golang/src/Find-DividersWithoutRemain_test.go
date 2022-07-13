package main

import "testing"

func TestMain(t *testing.T) {
        var test_res int
        test_res = Min([]int{17,78,87,3,18})
        if test_res != 3 {
                t.Error("Minimal value must be `3`, but here is:", test_res)
        }
}
