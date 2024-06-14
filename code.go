package main

import (
		"net"
		"bytes"
		"fmt"
)

var (
		// specify the range of trusted IP addresses
		start = net.ParseIP("192.168.1.0")
		end = net.ParseIP("192.168.1.199")
)

func checkIP(ip string) bool {

		//sanity check
		input := net.ParseIP(ip)
		if input.To4() == nil {
				fmt.Printf("%v is not a valid IPv4 address\n", input)
				return false
		}

		if bytes.Compare(input, start) >= 0 && bytes.Compare(input, end) <= 0 {
				fmt.Printf("%v is between %v and %v\n", input, start, end)
				return true
		}
		fmt.Printf("%v is NOT between %v and %v\n", input, start, end)
		return false
}

func main() {

		fmt.Println(checkIP("2.184.158.1"))
		fmt.Println(checkIP("1.2.3.4"))
		fmt.Println(checkIP("198.162.1.102"))
		fmt.Println(checkIP("1::16"))
}