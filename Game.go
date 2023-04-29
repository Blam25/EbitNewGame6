package main

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	var wg *sync.WaitGroup = &sync.WaitGroup{}

	for _, s := range Systems {
		wg.Add(1)
		go s(wg)
	}
	wg.Wait()
	//go TestSys()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	var wg *sync.WaitGroup = &sync.WaitGroup{}

	for _, s := range DrawSystems {
		wg.Add(1)
		go s(screen, wg)
	}
	wg.Wait()

	//go TestSys2(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

var Systems []func(wg *sync.WaitGroup)
var DrawSystems []func(screen *ebiten.Image, wg *sync.WaitGroup)

func init() {
	Systems = append(Systems, TestSys)
	DrawSystems = append(DrawSystems, TestSys2)
}

func TestSys(wg *sync.WaitGroup) {
	defer wg.Done()

	hmm := 5
	what := false

	Comps.Position.IterateWrite(func(i int, s *Position) {
		s.X = s.X + hmm
		print(s.X)
		if s.X > 500 {
			what = true
		}
	})

	if what {
		Comps.Position.IterateWrite(func(i int, s *Position) {
			print("hejsan")
		})
	}

	Comps.Position.IterateWrite(func(i int, s *Position) {
		if Comps.Image.Get(i) != nil {

		}
	})
}

func TestSys2(screen *ebiten.Image, wg *sync.WaitGroup) {
	defer wg.Done()

	Comps.Position.IterateRead(func(i int, s Position) {
		if a := Comps.Image.Get(i); a != nil {
			a.op.GeoM.Reset()
			a.op.GeoM.Translate(float64(s.X), float64(s.Y))
			screen.DrawImage(a.image, &a.op)
		}
	})
}
