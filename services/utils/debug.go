package utils

import "fmt"

func DebugPrintf(in interface{}) {
	fmt.Printf("%T %+v\n", in, in)
}
