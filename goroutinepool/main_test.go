package main

import (
	"fmt"
	"testing"
)

func Test_createPool(t *testing.T) {
	r_num := 3125
	var sum int
	for r_num != 0 {
		tmp := r_num % 10
		sum += tmp
		r_num /= 10
	}
	fmt.Println(sum)
}
