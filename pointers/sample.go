package pointers

import "fmt"

func PointerSample() {

	i, j := 10, 2000

	p := &i
	fmt.Println(*p)
	*p = 50
	fmt.Println(i)

	p = &j
	fmt.Println(*p)

}
