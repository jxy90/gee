package main

import (
	"fmt"
	example_first "github.com/jxy90/gee/gRPCuse/src/example.first"
)

func main() {
	//fmt.Println(1)
	NewPersonMessage()
}

func NewPersonMessage() {
	pm := example_first.PersonMessage{
		Id:       1234,
		IsAdult:  true,
		Name:     "Dave",
		LuckNums: []int32{1, 2, 3, 4, 5, 6, 7, 8},
	}
	fmt.Println(pm)
	fmt.Println(pm.GetId())
}
