/**
 * @Author xinwang_yao
 * @Date 2022/11/11 10:35
 **/

package main

import "C"
import "fmt"

func main1() {
	//str1 := 'abc' + '123'

	str2 := "abc" + "123"
	//str3 := 'abc' + "123"
	str4 := fmt.Sprintf("abc%d", 123)

	fmt.Println(str2)
	fmt.Println(str4)
	//字符串连接。除了以上两种连接方式，还有 strings.Join()、buffer.WriteString()等。
}

const (
	x = iota
	_ //跳过值
	y
	z = "zz"
	k
	g
	p = iota
)

func main2() {
	fmt.Println(x, y, z, k, p, g)
	fmt.Printf("x=%s\n", x)
	fmt.Printf("y=%s\n", y)
	fmt.Printf("z=%s\n", z)
	fmt.Printf("k=%s\n", k)
	fmt.Printf("p=%s\n", p)
	fmt.Printf("g=%s\n", g)
}

/**
    A. var x = nil
	B. var x interface{} = nil
    C. var x string = nil
    D. var x error = nil
	参考答案及解析：BD。知识点：nil 值。nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量。强调下 D 选项的 error 类型，它是一种内置接口类型，看下方贴出的源码就知道，所以 D 是对的。

type error interface {
    Error() string
}
 **/
func main() {
	//var x = nil
	//var x1 interface{} = nil
	//var x2 string = nil
	//var x3 error = nil
}
