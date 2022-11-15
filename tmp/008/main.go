/**
 * @Author xinwang_yao
 * @Date 2022/11/11 10:35
 **/

package main

import "C"
import "fmt"

func hello() []string {
	return nil
}

/**
A. nil

B. not nil

C. compilation error


	h := hello     B
	h := hello()   A
答案及解析：B。这道题目里面，是将 hello() 赋值给变量 h，而不是函数的返回值，所以输出 not nil。
*/

func main1() {
	h := hello
	//h := hello()
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

func GetValue() int {
	return 1
}

/**

考点：类型选择，类型选择的语法形如：i.(type)，其中 i 是接口，type 是固定关键字，需要注意的是，只有接口类型才可以使用类型选择。看下关于接口的文章。



Invalid type switch guard: i.(type) (non-interface type int on the left)
很显然 只有interface类型才可以使用类型选择
*/
//func main() {
//	i := GetValue()
//	switch i.(type) {
//	case int:
//		println("int")
//	case string:
//		println("string")
//	case interface{}:
//		println("interface")
//	default:
//		println("unknown")
//	}
//}
