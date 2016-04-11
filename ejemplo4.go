package main

import "fmt"

func main() {
	x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 99,17,
}
prom:=promedio(x)
	fmt.Println(prom)
}


	func promedio(vec []int) float64{
		total:=0.0
		for _, value := range vec {
			total+=float64(value)
		}
		return total/float64(len(vec))
	}
