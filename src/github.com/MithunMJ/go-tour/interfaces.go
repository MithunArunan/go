package main

import (
  "log"
)

type InterfaceName interface {
  FuncName() int
}

type StructName struct{
  S string
}

// Implicit interfaces -  Any type can implement an interface by implementing a method
func (structName *StructName)FuncName() int{
  log.Println("Function invoked: "+ structName.S)
  return 1
}

func main(){
  log.Println("Program Started")
  var InterfaceName InterfaceName
  structName := StructName{"Hello World!!"}
  InterfaceName = (&structName)
  InterfaceName.FuncName()
}
