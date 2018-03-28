package main

import "fmt"
import "encoding/json"

type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}


func main(){

  str := `{"page": 1, "fruits": ["apple", "peach"]}`
   res := response2{}
   json.Unmarshal([]byte(str), &res)
   fmt.Println(res)
}
