package main

import (
	//E "EbitNew6"
	"log"

	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/exp/maps"
)

var identifier int

type Remove struct {
	*Entity
}

func (s Remove) getid() int {
	return s.Entity.id
}

type Position struct {
	*Entity
	X int
	Y int
}

type Image struct {
	*Entity
	image *ebiten.Image
	op    ebiten.DrawImageOptions
}

func (s Image) getid() int {
	return s.Entity.id
}

func NewEntity() *Entity {
	new := &Entity{}
	new.id = identifier
	identifier++
	return new
}

type Entity struct {
	id int
}

func (s *Entity) Getid() int {
	return s.id
}

func NewComp[T validComp]() *Component[T] {
	new := &Component[T]{index: make(map[int]int)}
	return new

}

type Component[T validComp] struct {
	index    map[int]int
	theArray []T
	mu       sync.Mutex
}

type validComp interface {
	Getid() int
}

func (s *Component[T]) Add(object T) {

	s.index[object.Getid()] = len(s.theArray)
	s.theArray = append(s.theArray, object)

}

func (s *Component[T]) GetArrRead() []T {
	//for i, s := range s.theMap {
	//	f(i, s)
	//}
	/*for i, s := range s.theArray {
		z := *s
		f(i, z)
	}*/
	return s.theArray
}

func (s *Component[T]) IterateWrite(f func(entity int, object T)) {

	s.mu.Lock()
	//println("locked")
	for i, s := range s.theArray {
		f(i, s)
	}
	s.mu.Unlock()
	//println("Unlocked")
}

func (s *Component[T]) GetWrite(i int, f func(object T)) {

	s.mu.Lock()

	f(s.theArray[s.index[i]])

	s.mu.Unlock()
}

func (s *Component[T]) GetRead(id int) T {
	//return s.theMap[id]
	//z := s.theArray[id]
	return s.theArray[s.index[id]]
}

func (s *Component[T]) Remove(id int) {

	//get index of object to be removed
	index := s.index[id]

	//get object to be removed
	object := s.theArray[index]

	//delete id and index of said object from map
	delete(s.index, object.Getid())

	//set value of deleted index to the last object in array, deleting it
	s.theArray[index] = s.theArray[len(s.theArray)-1]

	//get id of moved index
	movedId := s.theArray[index].Getid()

	//set new index of moved object correctly in map
	s.index[movedId] = index

	//delete the last (now duplicated) object from the array
	s.theArray = s.theArray[:len(s.theArray)-1]
}

func (s *Component[T]) GetEntities() []int {
	return maps.Keys(s.index)
}

var PosMap map[int]*Position
var PosArr []*Position
var Positions *Component[*Position]
var Images *Component[*Image]

type Components struct {
	Position *Component[*Position]
	Image    *Component[*Image]
}

var Comps *Components = &Components{}
var Removed *Component[*Remove]

func init() {
	Removed = NewComp[*Remove]()
	Comps.Position = NewComp[*Position]()
	Comps.Image = NewComp[*Image]()

	var err error
	image1, _, err := ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	Ent1 := NewEntity()
	Comps.Position.Add(&Position{Ent1, 200, 200})
	Comps.Image.Add(&Image{Entity: Ent1, image: image1})

	Ent2 := NewEntity()
	Comps.Position.Add(&Position{Ent2, 100, 100})
	Comps.Image.Add(&Image{Entity: Ent2, image: image1})

	//Ent1.Add(&Position{}, PosMap)

}

func main() {
	ebiten.SetWindowSize(640, 480)
	//ebiten.SetVsyncEnabled(false)
	//ebiten.SetTPS(ebiten.SyncWithFPS)
	//ebiten.SetFPSMode(ebiten.)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
