package main

import "fmt"

func main() {
  fmt.Print("Ingrese un texto: ")
  var txt string
  fmt.Scanf("%s", &txt)

  fmt.Println(txt)

}
