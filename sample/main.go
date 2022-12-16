package main

import (
	"fmt"
	"go-training/pointers"
)

const Pi = 3.14

func main() {
	fizzbuzz()
	fizzbuzz2()
	pointers.PointerSample()
}

func fizzbuzz() {
	for i := 1; i <= 50; i++ {
		str := ""
		if i%3 == 0 {
			str += "FIZZ"
			if i%5 == 0 {
				str += "BUZZ"
			}
			fmt.Println(str, i)
		} else if i%5 == 0 {
			str += "BUZZ"
			fmt.Println(str, i)
		} else {
			fmt.Println(i)
		}

	}

}
func fizzbuzz2() {
	for i := 1; i <= 50; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FIZZBUZZ")
		case i%3 == 0:
			fmt.Println("FIZZ")
		case i%5 == 0:
			fmt.Println("BUZZ")
		default:
			fmt.Println(i)

		}
	}
}
