package main

import (
//  "log"
)

func main(){
  channel := make(chan bool)
  <-channel
}
