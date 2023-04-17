package main

import "log"

type A struct {
	ID uint32
}

func (a *A) GetID() uint32 {
	return a.ID
}

type entity interface {
	GetID() uint32
}

func ids(entities []entity) []uint32 {
	ids := make([]uint32, len(entities))
	for i, e := range entities {
		ids[i] = e.GetID()
	}
	return ids
}

func main() {
	a1, a2 := A{ID: 1}, A{ID: 2}
	entities := []A{a1, a2}
	log.Println(ids(entities))
	log.Println("hello")
}
