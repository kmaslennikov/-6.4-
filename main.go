package main

import "fmt"

func main() {
	var i10, i100, i1000 int64
  fmt.Println(i10)
	fmt.Println(i100)
	fmt.Println(i1000)

	for {
		if i1000 >= 1000 {
			break
		}
		i1000++

		if i100 >= 100 {
			continue
		}
		i100++

		if i10 >= 10 {
			continue
		}
		i10++
	}
	fmt.Println(i10)
	fmt.Println(i100)
	fmt.Println(i1000)
}
