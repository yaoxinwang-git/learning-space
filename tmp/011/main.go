/**
 * @Author xinwang_yao
 * @Date 2022/11/11 10:35
 **/

package main

import (
	"fmt"
)

/**
A. nil

B. not nil

C. compilation error

参考答案及解析：A。当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil。
*/
func main1() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

/**
A. runtime panic

B. 0

C. compilation error

参考答案及解析：B。删除 map 不存在的键值对时，不会报错，相当于没有任何作用；获取不存在的减值对时，返回值类型对应的零值，所以返回 0。
*/
func main() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}
