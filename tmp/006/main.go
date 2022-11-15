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

//& 取址运算符，* 指针解引用。
func main1() {
	//var p = &person{age: 10, name: "yxw"}
	p := new(person)
	p.name = "11"
	p.age = 10

	fmt.Printf("%T", p)

	fmt.Println(p.name)
	//fmt.Println((&p).name)
	fmt.Println((*p).name)
	//fmt.Println(p->name)
}

//类型定义 创建了新的类型 MyInt1
type MyInt1 int

//类型别名 本质还是int
type MyInt2 = int

func main() {
	var i int = 0
	//var i1 MyInt1 = i
	var i1 MyInt1 = (MyInt1(i))

	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}
