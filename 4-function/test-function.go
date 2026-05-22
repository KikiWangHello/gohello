package main

//函数定义格式
// func关键字 func1函数名（参数列表，参数） 返回值类型{
//}
func func1(a int, b string) int {
	println(a, b)
	c := a + 1
	return c
}

//多返回值的函数
// func 函数名（参数列表） （返回值类型1，类型2...）

func func2(a int, b bool) (int, bool) {
	return a, b
}

//返回值有名称的
//func 函数名（参数列表） (返回值变量1 类型1， 返回值变量2 类型2...)
func func3(a int, b int) (r1 int, r2 int) {
	r1 = a + 1
	r2 = b + 1
	return
}
func main() {
	println(func1(1, "hello"))
	a, b := func2(1, true)
	println(a, b)
	c, d := func3(1, 2)
	println(c, d)
}
