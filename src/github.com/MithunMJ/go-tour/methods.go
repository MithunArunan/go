package main

import (
//  "fmt"
  "log"
)

// Basic struct
type StructName struct {
  S string
}

//Method with pointer receivers, preferred over value receivers
func (structName *StructName) funcName() string {
  structName.S = "I'm changed now"
  return "Method invoked: "+structName.S
}

func main (){
    structName := StructName{"Hello"}
    log.Println(structName.S)
    log.Println(structName.funcName())
    log.Println(structName.S)
}
