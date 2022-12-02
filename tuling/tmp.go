/**
 * @Author xinwang_yao
 * @Date 2022/11/9 14:59
 **/

package main

import (
	"fmt"
)

//	func main() {
//		//var whatever = [5]int{1, 2, 3, 4, 5}
//		//for i, _ := range whatever {
//		//	//函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4.
//		//	y := i
//		//	defer func() { fmt.Println(y) }()
//		//}
//		//
//		//b := new(int)
//		//*b = 64
//		//
//		//fmt.Println(b)
//
//		type Player struct {
//			Name        string
//			HealthPoint int
//			MagicPoint  int
//		}
//		tank := new(Player)
//		tank.Name = "ms"
//		tank.HealthPoint = 300
//
//		//tank1 := &Player{}
//		var a *Player
//		a.MagicPoint = 10
//		a.HealthPoint = 10
//		a.Name = "test"
//		fmt.Printf("%T", a)
//	}
//
// 打印消息类型, 传入匿名结构体
func printMsgType(msg struct {
	id   int
	data string
}) {
	// 使用动词%T打印msg的类型
	fmt.Printf("%T\n, msg:%v", msg, msg)
}
func main2() {
	// 实例化一个匿名结构体
	msg := &struct { // 定义部分
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"hello",
	}
	printMsgType(*msg)
}

func main() {
	//

}
