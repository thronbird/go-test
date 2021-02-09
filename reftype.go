package main

type IP2 []byte

func (ip IP2) output() {
	print(ip)
}

func main() {
	var i = IP2{3, 4, 5, 9}
	i.output()
	var i2 = append(i, 5)
	i.output()
	i2.output()
}
