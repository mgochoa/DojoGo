
package main

import "fmt"

func main() {
  for i := 1; i <= 100; i++ {
	x:=""
    if i % 3 == 0 {
      x+="go "
    }
    if i % 5 == 0 {
      x+="team"
    }
    fmt.Println(i, x)
  }
}
