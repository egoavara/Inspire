package main

import (
	"sync/atomic"
	"fmt"
)

func main() {
	var a uint32 = 4
	fmt.Println(atomic.CompareAndSwapUint32(&a, 4, 0))
	fmt.Println(a)
}
