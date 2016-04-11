package main

import "fmt"

func main() {
  fmt.Print("Ingrese un texto: ")
  var txt string
  fmt.Scanf("%s", &txt)
  x:=""
  for i:= 0; i< len(txt) ; i++{
  	x=txt[0:i]
  	fmt.Println(x)
  }
  for i:= len(txt); i>=0 ; i--{
  	x=txt[0:i]
  	fmt.Println(x)
  }
}
