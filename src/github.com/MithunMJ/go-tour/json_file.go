package main

import "io/ioutil"
import "encoding/json"
import "log"

// Converts JSON file to a interface
func load_json(file string)(interface{}) {
  file_stream, _ := ioutil.ReadFile(file)
  var data interface{}
  err := json.Unmarshal(file_stream, &data)
  if err != nil {
    log.Print("Error in unmarshalling JSON")
  }
  return data
}

func main(){
  data := load_json("config.json")
  dat := data.(map[string]interface{})
  d := dat["key"].([]interface{})
  log.Print(d)
  //data.(map([]map))
}
