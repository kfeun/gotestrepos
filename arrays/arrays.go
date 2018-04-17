package main

import (
	"fmt"
)

func main() {

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64
	// old cycle
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }
	for _, value := range x {
		total += value
	}
	// Одиночный символ подчеркивания _ используется,
	// чтобы сказать компилятору, что переменная нам не нужна
	// (в данном случае нам не нужна переменная итератора).
	fmt.Println(total / float64(len(x)))

	y := [5]float64{
		98,
		93,
		77,
	}
	fmt.Println(y)
	slices()
	mapFunc()
}

func slices() {
	fmt.Println("slices")

	x := make([]float64, 5, 10)
	fmt.Println(x)

	arr := [5]float64{1, 2, 3, 4, 5}
	y := arr[1:4]
	fmt.Println(y)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}

func mapFunc() {
	fmt.Println("map")
	// карта string-ов для int-ов
	x := make(map[string]int)
	x["key"] = 10
	x["Un"] = 10
	fmt.Println(x["key"])
	fmt.Println(len(x))
	delete(x, "key")
	fmt.Println(x)

	if name, ok := x["Un"]; ok {
		fmt.Println(name, ok)
	}

	elements := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements)

		elements2 := map[string]map[string]string{
				"H": map[string]string{
				"name":"Hydrogen",
				"state":"gas",
			},
			"He": map[string]string{
				"name":"Helium",
				"state":"gas",
			},
			"Li": map[string]string{
				"name":"Lithium",
				"state":"solid",
			},
			"Be": map[string]string{
				"name":"Beryllium",
				"state":"solid",
			},
			"B":  map[string]string{
				"name":"Boron",
				"state":"solid",
			},
			"C":  map[string]string{
				"name":"Carbon",
				"state":"solid",
			},
			"N":  map[string]string{
				"name":"Nitrogen",
				"state":"gas",
			},
			"O":  map[string]string{
				"name":"Oxygen",
				"state":"gas",
			},
			"F":  map[string]string{
				"name":"Fluorine",
				"state":"gas",
			},
			"Ne":  map[string]string{
				"name":"Neon",
				"state":"gas",
			},
		}
	
		if el, ok := elements2["Li"]; ok {    
			fmt.Println(el["name"], el["state"])
		}

}
