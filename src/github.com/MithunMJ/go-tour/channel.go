package main

import (
  "log"
)

func hello(channel chan int, num int) {
  channel <- num
}


func main(){
  channel := make(chan int, 2)
  go hello(channel,1)
  go hello(channel,2)
  go hello(channel,3)
  go hello(channel,4)
  for i := range channel{
    log.Println(i)
  }
}
