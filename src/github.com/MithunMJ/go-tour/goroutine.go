package main

import (
  "log"
  "fmt"
)

func funcName(){
  log.Println("A Goroutine is a lightweight thread managed by the Go runtime.")
}

func main(){
  go funcName()
  fmt.Scanln()
}
