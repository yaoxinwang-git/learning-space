// package main
//
// import "fmt"
//
// // 支持返回值 命名 ，默认值为类型零值,命名返回参数可看做与形参类似的局部变量,由return隐式返回
//
//	func f1() (names []string, m map[string]int, num int) {
//		m = make(map[string]int)
//		m["k1"] = 2
//
//		return
//	}
//
//	func main() {
//		a, b, c := f1()
//		fmt.Println(a, b, c)
//	}
//package main
//
//import (
//	"fmt"
//)
//
//// 创建一个玩家生成器, 输入名称, 输出生成器
//func playerGen(name string) func() (string, int) {
//	// 血量一直为150
//	hp := 150
//	// 返回创建的闭包
//	return func() (string, int) {
//		// 将变量引用到闭包中
//		return name, hp
//	}
//}
//
//// 创建一个玩家生成器, 输入名称, 输出生成器
//func playerGen1() func(string) (string, int) {
//	// 血量一直为150
//	hp := 150
//	// 返回创建的闭包
//	return func(name string) (string, int) {
//		// 将变量引用到闭包中
//		return name, hp
//	}
//}
//func main() {
//	// 创建一个玩家生成器
//	generator := playerGen("ms")
//	// 返回玩家的名字和血量
//	name, hp := generator()
//	// 打印值
//	fmt.Println(name, hp)
//
//	generator1 := playerGen1()
//	name1, hp1 := generator1("ms")
//	// 打印值
//	fmt.Println(name1, hp1)
//}

package main

import (
	"fmt"
	"log"
	"time"
)

func main1() {
	//var whatever = [5]int{1, 2, 3, 4, 5}
	//
	//for i := range whatever {
	//	defer fmt.Println(i)
	//}

	start := time.Now()
	log.Printf("开始时间为：%v", start)
	defer log.Printf("时间差：%v", time.Since(start)) // Now()此时已经copy进去了
	//不受这3秒睡眠的影响
	time.Sleep(20 * time.Second)

	log.Printf("函数结束")
}

//	func main() {
//		//start := time.Now()
//		//log.Printf("开始时间为：%v", start)
//		//defer func() {
//		//	log.Printf("开始调用defer")
//		//	log.Printf("时间差：%v", time.Since(start))
//		//	log.Printf("结束调用defer")
//		//}()
//		//time.Sleep(3 * time.Second)
//		//
//		//log.Printf("函数结束")
//
//		var whatever = [5]int{1, 2, 3, 4, 5}
//		for i, _ := range whatever {
//			i := i
//			defer func() { fmt.Println(i) }()
//		}
//	}
func test() {
	defer func() {
		// defer panic 会打印
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}

func main() {
	test()
}
