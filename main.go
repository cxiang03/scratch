package main

import (
	"log"
	"sort"
)

func main() {
	s := []string{}
	s = nil
	sort.Strings(s)
	log.Println(s)
}
