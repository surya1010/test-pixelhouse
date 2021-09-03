package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	for {
		fmt.Print("silahkan ketik angka : ")

		var input int

		fmt.Scanf("%d", &input)
		if input < 1 {
			fmt.Println("hanya bisa di isi angka")
			break
		}
		var data []string
		for i := 0; i < input+1; i++ {
			data = append(data, strconv.Itoa(rumus(i)))
		}

		fmt.Println(strings.Join(data, ", "))
	}

}

//n(n+1)/2 + 1
func rumus(n int) int {
	var quadrant int
	quadrant = n + 1

	b := (n * quadrant / 2) + 1
	return b
}
