package main

import "fmt"

//	func main() {
//		/*var mapLit map[string]int
//		var mapAssigned map[string]int
//		mapLit = map[string]int{"one": 1, "two": 2}
//		mapAssigned = mapLit
//		//mapAssigned 是 mapList 的引用，对 mapAssigned 的修改也会影响到 mapList 的值。
//		mapAssigned["two"] = 3
//		fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
//		fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
//		fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
//		*/
//		// 创建一个int到int的映射
//		m := make(map[int]int)
//		// 开启一段并发代码
//		go func() {
//			// 不停地对map进行写入
//			for {
//				m[1] = 1
//			}
//		}()
//		// 开启一段并发代码
//		go func() {
//			// 不停地对map进行读取
//			for {
//				_ = m[1]
//			}
//		}()
//		// 无限循环, 让并发程序在后台执行
//		for {
//		}
//	}
//
//	func main() {
//		//sync.Map 不能使用 make 创建
//		var scene sync.Map
//		// 将键值对保存到sync.Map
//		//sync.Map 将键和值以 interface{} 类型进行保存。
//		scene.Store("greece", 97)
//		scene.Store("london", 100)
//		scene.Store("egypt", 200)
//		// 从sync.Map中根据键取值
//		fmt.Println(scene.Load("london"))
//		// 根据键删除对应的键值对
//		scene.Delete("london")
//		// 遍历所有sync.Map中的键值对
//		//遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}，每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回。
//		scene.Range(func(k, v interface{}) bool {
//			fmt.Println("iterate:", k, v)
//			return true
//		})
//
// //	}
//
//	func main() {
//		var arr []int
//		var num *int
//		fmt.Printf("%p\n", arr)
//		fmt.Printf("%p", num)
//
//		var p *struct{}
//		fmt.Println(unsafe.Sizeof(p)) // 8
//		var s []int
//		fmt.Println(unsafe.Sizeof(s)) // 24
//		var m map[int]bool
//		fmt.Println(unsafe.Sizeof(m)) // 8
//		var c chan string
//		fmt.Println(unsafe.Sizeof(c)) // 8
//		var f func()
//		fmt.Println(unsafe.Sizeof(f)) // 8
//		var i interface{}
//		fmt.Println(unsafe.Sizeof(i)) // 16
//	}
//func main() {
//	var a int
//	for x := 0; x < 10; x++ {
//		for y := 0; y < 10; y++ {
//			if y == 2 {
//				// 跳转到标签
//				goto breakHere
//			}
//		}
//	}
//	// 手动返回, 避免执行进入标签
//	a = 10
//	fmt.Println("doing")
//	fmt.Println(a)
//
//	// 标签
//breakHere:
//	fmt.Println("done")
//}

func main() {
	str := "ms的go教程"
	//因为一个字符串是 Unicode 编码的字符（或称之为 rune ）集合
	//char 实际类型是 rune 类型
	for pos, char := range str {
		fmt.Println(pos, char)
	}
}
