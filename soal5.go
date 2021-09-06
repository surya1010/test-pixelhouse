package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Printf("silahkan input alphhabet, contoh : A-B-C atau a-b-c\n")

	for {
		models := map[int]interface{}{
			0: "A",
			1: map[string][]string{
				"B": {"D", "E"},
			},
			2: map[string][]string{
				"C": {"F", "G", "H"},
			},
		}

		fmt.Printf("silahkan input alphhabet : ")

		var input string

		fmt.Scanf("%s", &input)

		arrayInput := strings.Split(strings.ToUpper(input), "-")

		var selectionRoot bool
		var results []string

		if reflect.ValueOf(models[0]).Kind() == reflect.String && arrayInput[0] == "A" {
			selectionRoot = true

		}

		for k := 1; k <= 2; k++ {
			if reflect.ValueOf(models[k]).Kind() == reflect.Map {
				isMap, ok := models[k].(map[string][]string)
				if !ok {

					break
				}
				if len(isMap) > 0 {
					for node, parentNodes := range models[k].(map[string][]string) {

						if node == arrayInput[1] && selectionRoot {

							if len(parentNodes) > 0 {
								for _, parent := range parentNodes {
									results = append(results, parent)
								}

							}
						} else {
							continue
						}

					}
				}

			}
		}
		if len(results) > 0 {
			fmt.Println(strings.Join(results, "-"))
		} else {
			fmt.Println("Tidak ada child")
		}

	}

}
