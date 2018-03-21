package main

import "fmt"

type IPAddr [4]byte

func (ipaddr *IPAddr)String() string{
  var fmt_ip string
	for ip := range ipaddr{
    fmt.Print(ip)
  }
  return fmt_ip+"hello"
}


func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for _, ip := range hosts {
		fmt.Println(&ip)
	}
}
