package main

import "fmt"

func main() {

	//var a = [3]int{1, 2, 3}
	////a[1:2] 生成了一个新的切片
	//var b = a[1:3]
	//c := make([]int, 0)
	//c = b
	//fmt.Println(&b[0], &b[1])
	////fmt.Println(a, a[1:2])
	//fmt.Println(&a[1], &a[2])
	//
	//fmt.Println(&c[0], &c[1])

	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6]
	//这打印出来长度为2
	fmt.Printf("myslice为 %d, 其长度为: %d\n", myslice, len(myslice))

	fmt.Println(cap(myslice))
	myslice = myslice[:cap(myslice)]
	//为什么 myslice 的长度为2，却能访问到第四个元素
	fmt.Printf("myslice的第四个元素为: %d\n", myslice[3])
	fmt.Printf("myslice的第四个元素为: %d\n", myslice[4])
	fmt.Printf("myslice的第四个元素为: %d\n", myslice[5])
	fmt.Printf("myslice的第四个元素为: %d\n", myslice[6])
	fmt.Printf("myslice的第四个元素为: %d\n", myslice[7])

}
