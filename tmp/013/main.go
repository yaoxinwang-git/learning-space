/**
 * @Author xinwang_yao
 * @Date 2022/11/25 11:44
 **/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Book struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

func main() {

	file, err := ioutil.ReadFile("D:\\workspace\\src\\learning-space\\tmp\\013\\io.conf")
	if err != nil {
		fmt.Printf("读取文件出错%v\n", err)
	}

	configs := make([]Book, 0)
	err = json.Unmarshal(file, &configs)
	if err != nil {
		fmt.Printf("序列化出错%v\n", err)

	}
	fmt.Printf("books: %v\n", configs)

}
