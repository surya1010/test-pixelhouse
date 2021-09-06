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
		var selectionParent bool

		if reflect.ValueOf(models[0]).Kind() == reflect.String && arrayInput[0] == "A" {
			selectionRoot = true

		}

		for k := 1; k <= 2; k++ {
			if reflect.ValueOf(models[k]).Kind() == reflect.Map {
				isMap, ok := models[k].(map[string][]string)
				if !ok {
					selectionParent = false
					break
				}
				if len(isMap) > 0 {
					for node, parentNodes := range models[k].(map[string][]string) {

						if node == arrayInput[1] {

							if len(parentNodes) > 0 {
								var childrenCheck = 0

								for _, element := range arrayInput[2:] {
									hasElement := contains(parentNodes, element)
									if !hasElement {
										childrenCheck = 0
										break
									}

									for _, parent := range parentNodes {

										if parent == element {
											childrenCheck++
										}

									}

								}
								if childrenCheck == 1 {
									selectionParent = true

								}

							}
						} else {
							continue
						}

					}
				}

			}
		}

		fmt.Println(selectionRoot && selectionParent)
	}

}

func contains(s []string, e string) bool {

	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
