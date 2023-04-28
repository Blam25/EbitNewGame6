package main

import (
	//E "EbitNew6"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/exp/maps"
)

var identifier int

type Position struct {
	X int
	Y int
}

type Entity struct {
	Id int
}

func NewComp[T any]() *Component[T] {
	new := &Component[T]{theMap: make(map[int]*T)}
	return new

}

type Component[T any] struct {
	theMap map[int]*T
}

func (s *Component[T]) Add(entity *Entity, object *T) {
	s.theMap[entity.Id] = object
}

func (s *Component[T]) Iterate(f func(object *T)) {
	for _, s := range s.theMap {
		f(s)
	}
}

func (s *Component[T]) Get(id int) *T {
	return s.theMap[id]
}

func (s *Component[T]) Remove(id int) {
	delete(s.theMap, id)
}

func (s *Component[T]) GetEntities() []int {
	return maps.Keys(s.theMap)
}

func NewEntity() *Entity {
	new := &Entity{}
	new.Id = identifier
	identifier++
	return new
}

var PosMap map[int]*Position
var PosArr []*Position
var Positions *Component[Position]

func init() {

	Positions = NewComp[Position]()

	print(Positions)

	Ent1 := NewEntity()
	Positions.Add(Ent1, &Position{200, 200})
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
