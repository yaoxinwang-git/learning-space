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

func main() {
	//[keytype] 和 valuetype 之间允许有空格。
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}
