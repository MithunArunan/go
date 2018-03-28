package main

//import "io/ioutil"
import "encoding/json"
import "log"

type Data struct{
  Key string `json:"key"`
  Key2 string `json:"key2"`
}

func main(){
  data := Data{}
  //file := "config.json"
  //file_stream, _ := ioutil.ReadFile(file)
  str := `{"key": "hi", "fruits": ["apple", "peach"]}`
  err := json.Unmarshal([]byte(str), &data)
  if err != nil {
    log.Print("Error in unmarshalling JSON")
  }
  log.Print(data.Key)
}
