package main

import "fmt"


func main() {

per:=Persona{"Manuel"}
	fmt.Println(per.correr())
}


type Persona struct{
	nombre string
}

func (per *Persona) correr() string {
  x:="Corra, "
  x+=per.nombre
	x+=", Corra!!!!!"
	return x
}
