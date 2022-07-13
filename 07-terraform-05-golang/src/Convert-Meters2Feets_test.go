package main

import "testing"

func TestMain(t *testing.T) {
        var test_res float32
        test_res = convert_meters_to_feets(77)
        if test_res != 252.62466 {
                t.Error("Value must be `252.62466`, but here is:", test_res)
        }
}
