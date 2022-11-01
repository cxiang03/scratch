package main

import "nothing.com/scratch/pb"

type Slot struct {
	Name string
}

// Pb returns the protobuf value from the Slot
// as its default schema
// mostly for list API or normal case API.
func (s *Slot) Pb() *pb.Slot {
	return &pb.Slot{
		Name: s.Name,
	}
}

// PbDetail returns the protobuf value from the Slot
// as its detail schema
// mostly for detail item API or some special case.
func (s *Slot) PbDetail() *pb.Slot {
	rst := s.Pb()
	_ = rst // do some extra work
	return rst
}

// NewSlotFrom returns a new Slot from the protobuf value.
func NewSlotFrom(pb *pb.Slot) *Slot {
	return &Slot{
		Name: pb.Name,
	}
}

// Slots is a slice of Slot
// it is not recommended do slice aliasing, since it is not easy to determine when to use pointer or not
// let it be a slice is better for understanding and easy to code.
type Slots []Slot

// Pb returns the protobuf value from the Slots
// but in some case, we still need to decode into a slice of protobuf value
// just return a slice of protobuf value directly.
func (s Slots) Pb() []*pb.Slot {
	pb := make([]*pb.Slot, len(s))
	for i := range s {
		pb[i] = s[i].Pb()
	}
	return pb
}

// PbDetail returns the protobuf value from the Slots.
func (s Slots) PbDetail() []*pb.Slot {
	pb := make([]*pb.Slot, len(s))
	for i := range s {
		pb[i] = s[i].PbDetail() // Convention, Convention, Convention
	}
	return pb
}

// NewSlotFrom returns a new Slot from the protobuf value.
// following the same pattern as Slot.Pb().
func NewSlotsFrom(pb []*pb.Slot) Slots {
	slots := make(Slots, 0, len(pb))
	for _, slot := range pb {
		slots = append(slots, *NewSlotFrom(slot))
	}
	return slots
}
