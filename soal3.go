package main

import (
	"fmt"
	"reflect"
)

type (
	Tree struct {
		Root   int
		Node   int
		Parent []int
	}
)

func main() {

	data := map[int]interface{}{
		0: "A",
		1: map[string][]string{
			"B": {"D", "E"},
		},
		2: map[string][]string{
			"C": {"F", "G", "H"},
		},
	}

	var recursiveAlpha func(map[int]interface{}, int)

	recursiveAlpha = func(models map[int]interface{}, variable int) {

		if variable > 2 {

			fmt.Println("**/")
			return
		} else {
			if variable == 0 {
				fmt.Println("* ", models[variable])
			} else {

				if reflect.ValueOf(models[variable]).Kind() == reflect.Map {

					isMap, ok := models[variable].(map[string][]string)
					if !ok {
						return
					}
					if len(isMap) > 0 {
						for node, parentNodes := range models[variable].(map[string][]string) {
							fmt.Printf("*  -%s\n", node)
							if len(parentNodes) > 0 {
								for _, parent := range parentNodes {
									fmt.Printf("*  --%s\n", parent)
								}
							}

						}
					}

				}

			}

			recursiveAlpha(models, variable+1)
		}

	}

	fmt.Println("/**")
	recursiveAlpha(data, 0)

}

// func toCharStrArr(i int) string {
// 	return arr[i]
// }
