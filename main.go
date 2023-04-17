package main

import "log"

func isNil(a interface{}) bool {
	return a == nil
}

func main() {
	var a []uint32
	log.Println(a)

	i := isNil(a)
	log.Println(i)
}
