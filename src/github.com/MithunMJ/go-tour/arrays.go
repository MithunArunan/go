package main

import (
//  "fmt"
  "log"
)

func main (){
    structName := StructName{"Hello"}
    log.Println(structName.S)
    log.Println(structName.funcName())
    log.Println(structName.S)
}
