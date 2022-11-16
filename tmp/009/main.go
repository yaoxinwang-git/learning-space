/**
 * @Author xinwang_yao
 * @Date 2022/11/11 10:35
 **/

package main

import (
	"fmt"
)

type person struct {
	name string
}

//打印一个 map 中不存在的值时，返回元素类型的零值。这个例子中，m 的类型是 map[person]int，因为 m 中不存在 p，所以打印 int 类型的零值，即 0。
func main2() {
	//[keytype] 和 valuetype 之间允许有空格。
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}

func hello(num ...int) {
	num[0] = 18
}

//map、slice、chan、指针、interface默认以引用的方式传递，修改会实际影响原切片
func main() {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Printf("%T\n", i) //i的类型是int切片
	fmt.Println(i[0])
}
