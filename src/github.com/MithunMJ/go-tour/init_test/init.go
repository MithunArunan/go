package init_test

import "fmt"
import "log"
import "math/rand"

var a string

func init(){
  a = fmt.Sprint(rand.Float64()*2)

}

func Hello(){
  log.Printf(a)
}
