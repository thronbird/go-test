package main

type plant struct {
	size int
}

type fruit struct {
	p      plant
	flavor string
}

type fruit2 struct {
	plant
	flavor string
}

func (*plant) grow() {
	println("2221")
}

func (*fruit) grow() {
	println("22212")
}

func main() {
	a := fruit{
		plant{3}, "sssss",
	}
	println(a.p.size)
	a.grow()
	a.p.grow()

	b := fruit2{
		plant{3}, "sssss",
	}
	println(b.size)
	println(b.plant.size)
	b.grow()
	b.plant.grow()
}
