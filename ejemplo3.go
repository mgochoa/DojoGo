package main

import "fmt"

func main() {
	x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 99,17,
}

	minimo:=x[0]
	for _, value := range x {
		if value < minimo {
			minimo=value
		}
	}
	fmt.Println(minimo)
}
