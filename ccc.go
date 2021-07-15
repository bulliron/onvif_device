package main

import (
	"fmt"
	"net"
	"reflect"
)
// net.Interfaces : Test
func NetTest(){
	fmt.Println("\n- netTest() - ")
	//fmt.Println("Hi")
	l, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error")
	}
	//fmt.Println(l)
	t := reflect.TypeOf(l)
	//fmt.Println("t.Align() : ",t.Align())
	// fmt.Println("t.FieldAlign() : ",t.FieldAlign())
	// fmt.Println("t.NumMethod() : ",t.NumMethod())
	// fmt.Println("t.Name() : ",t.Name()) // 
	// fmt.Println("t.PkgPath() : ",t.PkgPath())
	// fmt.Println("t.Size() : ",t.Size())
	fmt.Println("t.Kind() : ", t.Kind())
	fmt.Println("t.String() : ", t.String())
	for _, c :=range(l) {
		fmt.Println("net.Interface.Index = ", c.Index)
		fmt.Println("net.Interface.MTU = ", c.MTU)
		fmt.Println("net.Interface.Name = ", c.Name)
		fmt.Println("net.Interface.HardwareAddr = ", c.HardwareAddr)
		fmt.Println("net.Interface.Flags = ", c.Flags)
	}
}
func sliceTest() {
	var s []int = []int{1,2,3} // slice에 literal 대입
	fmt.Println("\n- sliceTest() -")
	fmt.Println("s = ",s)
	t := make([]int,0,3) // slice type, length, capability 
	fmt.Println("t = ",t)
}
// func structTest() {
// 	s := &ttt{u : "111", t: "dddd"}
// 	fmt.Println(s)
// }





