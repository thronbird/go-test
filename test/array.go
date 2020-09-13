package main

func main() {
	arr := [...]int{1, 2, 3, 5}
	print(arr[3])

	// 声明一个需要 8 MB 的数组
	var array [1e6]int
	// 不使用指针会栈上分配8m的内存，浪费
	foo(array)
	foo2(&array)
}

func foo(arr [1e6]int) {
	print(arr[1])
}

func foo2(arr *[1e6]int) {
	print(arr[1])
}
