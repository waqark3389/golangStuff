package main

import (
	"fmt"
	"net"
)

func main() {

	IP := "10.4.26.8"

	mdhNet := "10.4.0.0/16"

	ipAddressByte := net.ParseIP(IP)
	_, network, err := net.ParseCIDR(mdhNet)

	fmt.Println(network.Contains(ipAddressByte))

	if err != nil {
		fmt.Println("ip parse error" + err.Error())
		return
	}

	//fmt.Println(ipnetA)

	if network.Contains(ipAddressByte) == true {
		fmt.Println("IP: " + IP + " exists in " + mdhNet)
	} else {
		fmt.Println("IP: " + IP + " does not exist in " + mdhNet)
	}
	return

}
