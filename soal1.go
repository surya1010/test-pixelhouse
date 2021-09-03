package main

import (
	"fmt"
)

func main() {

	a := 1
	for i := 0; i < 100; i++ {

		if i > 0 {
			a = a + i
			if i == 20 || i == 50 {
				fmt.Println("dipotong", i, "jadi", a)
			}

		}
	}

}
