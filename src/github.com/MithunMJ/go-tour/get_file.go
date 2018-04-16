package main

import (
//  "log"
//  "net/url"
  "github.com/hashicorp/go-getter"
)

func main(){
	file_url_path := "https://hello.com/hello"
	//file_url,_ := url.Parse(file_url_path)

	// Build the client
	client := &getter.Client{
			Src:  file_url_path,
			Dst:  "hello.mp3",
			// Pwd:
			Mode: getter.ClientModeFile,
	}
	client.Get()
}
