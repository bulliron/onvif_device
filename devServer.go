package main

import (
	"fmt"
	"net"
	//"reflect"
)
func myServer(){
	discoveryMulticastAddr, err := net.ResolveUDPAddr("udp","239.255.255.250:3702")
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("net.ResolveUDPAddr : ",discoveryMulticastAddr)
	}
	//fmt.Println(discoveryMulticastAddr)
	netInterface, err :=net.InterfaceByName("en0")
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("net.InterfaceByName : ", netInterface)
	}
	discoveryConn, err := net.ListenMulticastUDP("udp", netInterface, discoveryMulticastAddr)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("net.ListenMulticastUDP : ", discoveryConn)
	}
	for {
		b := make([]byte, 1024)
		n, err := discoveryConn.Read(b)
		if err != nil {
			fmt.Println("Error : ", err)
		} else {
			fmt.Println("Read : ", n)
		}
	}
}
