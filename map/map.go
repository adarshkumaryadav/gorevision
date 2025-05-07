package main

import "fmt"

func main() {
	// this is a empty map
	// var myMap map[string]string
	var myMap = make(map[string]string)
	fmt.Println(len(myMap))  // 0 len
	fmt.Println(myMap)       // map[]
	myMap["key1"] = "value1" // will panic as map is empty
	// after initializing at line no. above will not panic
	fmt.Println(myMap)
}
