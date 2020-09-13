package main

type aaa struct {
	ab bbb
}

type bbb struct {
}

func (bbb) b1() {

}

func (a aaa) b1() interface{} {
	x := [5]int{1, 2, 3, 4, 5}
	print(x[2])
	return nil
}

func main() {
	var x = aaa{}
	print(x.b1())
}
