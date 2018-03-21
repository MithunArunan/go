// Basic struct
package main

import (
  "fmt"
)

type StructName struct {
  S string
}


func main (){
    structName := StructName{"Hello"}
    fmt.Println(structName.S)
}
