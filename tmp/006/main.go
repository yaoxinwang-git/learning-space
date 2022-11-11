/**
 * @Author xinwang_yao
 * @Date 2022/11/11 10:35
 **/

package main

import "fmt"

type person struct {
	age  int
	name string
}

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	fmt.Printf("%T", sn1)
}
