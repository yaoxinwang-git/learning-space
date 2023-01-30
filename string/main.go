/**
 * @Author xinwang_yao
 * @Date 2023/1/10 14:55
 **/

package string

import (
	"bytes"
	"strings"
)

//strings.Join ≈ strings.Builder > bytes.Buffer > "+" > fmt.Sprintf

/*
*
1.使用+ 操作符进行拼接时，会对字符串进行遍历，计算并开辟一个新的空间来存储原来的两个字符串。
2.fmt.Sprintf 参数中有个any类型 本质为interface{}  由于采用了接口参数，必须要用反射获取值，因此有性能损耗。
*/
func main() {

	//fmt.Sprintf()
	strings.Builder{}.String()
}

func exampleBytesBuffer() {
	// bytes.Buffer的0值可以直接使用
	var buff bytes.Buffer

	// 向buff中写入字符/字符串
	buff.Write([]byte("Hello"))
	buff.WriteByte(' ')
	buff.WriteString("World")

	// String() 方法获得拼接的字符串
	buff.String() // "Hello World"
}
func exampleStringsBuilder() {
	// strings.Builder的0值可以直接使用
	var builder strings.Builder

	// 向builder中写入字符/字符串
	builder.Write([]byte("Hello"))
	builder.WriteByte(' ')
	builder.WriteString("World")

	// String() 方法获得拼接的字符串
	builder.String() // "Hello World"
}
