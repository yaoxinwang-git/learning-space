/**
 * @Author xinwang_yao
 * @Date 2022/11/16 15:36
 **/

package main

import "fmt"

//func main1() {
//	//http://127.0.0.1:8000/go
//	// 单独写回调函数
//	http.HandleFunc("/go", myHandler)
//	// addr：监听的地址
//	// handler：回调函数
//	http.ListenAndServe("127.0.0.1:8000", nil)
//}
//
//// handler函数
//func myHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println(r.RemoteAddr, "连接成功")
//	// 请求方式：GET POST DELETE PUT UPDATE
//	fmt.Println("method:", r.Method)
//	// /go
//	fmt.Println("url:", r.URL.Path)
//	fmt.Println("header:", r.Header)
//	fmt.Println("body:", r.Body)
//	// 回复
//	w.Write([]byte("你好，ms的go教程"))
//}

//a 的类型是 int，b 的类型是 float，两个不同类型的数值不能相加，编译报错
//func main2() {
//	a := 5
//	b := 8.1
//	fmt.Println(a + b)
//}

/**

B。知识点：操作符 [i:j]。基于数组（切片）可以使用操作符 [i,j] 创建新的切片，从索引 i，到索引 j 结束，截取已有数组（切片）的任意部分，返回新的切片，
新切片的值包含原数组（切片）的 i 索引的值，但是不包含 j 索引的值。i、j 都是可选的，i 如果省略，默认是 0，j 如果省略，默认是原数组（切片）的长度。
i、j 都不能超过这个长度值。
假如底层数组的大小为 k，截取之后获得的切片的长度和容量的计算方法：长度：j-i，容量：k-i。
截取操作符还可以有第三个参数，形如 [i:j:k]，第三个参数 k 用来限制新切片的容量，但不能超过原数组（切片）的底层数组大小。截取获得的切片的长度和容量分别是：j-i、k-i。
所以例子中，切片 t 为 [4]，长度和容量都是 1。
*/

func main3() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])
	fmt.Println(len(t))
	fmt.Println(cap(t))

}

/**
A. compilation error

B. equal

C. not equal

参考答案及解析：A。Go 中的数组是值类型，可比较，另外一方面，数组的长度也是数组类型的组成部分，所以 a 和 b 是不同的类型，是不能比较的，所以编译错误。

*/
//func main() {
//	a := [2]int{5, 6}
//	b := [3]int{5, 6}
//	if a == b {
//		fmt.Println("equal")
//	} else {
//		fmt.Println("not equal")
//	}
//}
