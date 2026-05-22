package main

import "fmt"

func main() {
	// 不赋值，使用默认未0
	var a int = 10
	fmt.Println(a)                               //换行输出
	fmt.Printf("a = %d, type of a = %T\n", a, a) //格式化输出

	var b = "asd"
	fmt.Printf("type of b = %T", b)

	c := "hello"
	fmt.Printf("type of c = %T\n", c)

	d := 3.14
	fmt.Printf("d = %f\n", d)
	fmt.Printf("d = %T\n", d)
}
