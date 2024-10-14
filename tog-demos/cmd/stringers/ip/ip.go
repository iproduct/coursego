package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	addr := make([]any, 4)
	for i, octet := range ip {
		addr[i] = octet
	}
	return fmt.Sprintf("%d.%d.%d.%d", addr...)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
