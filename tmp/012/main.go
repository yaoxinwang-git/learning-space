/**
 * @Author xinwang_yao
 * @Date 2022/11/11 10:35
 **/

package main

import (
	"fmt"
)

/**
A. -5 +5

B. +5 +5

C. 0 0

参考答案及解析：A。%d表示输出十进制数字，+表示输出数值的符号。这里不表示取反。

*/
func main1() {
	i := -5
	j := +5
	fmt.Printf("%+d %+d", i, j)
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
