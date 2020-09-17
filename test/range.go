package main

func main() {
	/*x := []int{2,3,4,7,8}

	for i,y :=  range x{
		print(i)
		print(y)
		print("xxxxxxx")
	}*/

	z := map[string]string{"1": "uu", "ii": "pp"}
	for key, val := range z {
		println(key)
		println(val)
	}

	print(z["dwssfa"])
}
