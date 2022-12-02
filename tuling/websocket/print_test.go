/**
 * @Author xinwang_yao
 * @Date 2022/11/28 11:48
 **/

package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFPrint(t *testing.T) {
	fmt.Fprintln(os.Stdout, "向标准输出写入字符串")
}

func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
