package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {

	TestSys()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	TestSys2(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func TestSys() {

	hmm := 5
	what := false

	Comps.Position.Iterate(func(i int, s *Position) {
		s.X = s.X + hmm
		print(s.X)
		if s.X > 500 {
			what = true
		}
	})

	if what {
		Comps.Position.Iterate(func(i int, s *Position) {
			print("hejsan")
		})
	}

	Comps.Position.Iterate(func(i int, s *Position) {
		if Comps.Image.Get(i) != nil {

		}
	})
}

func TestSys2(screen *ebiten.Image) {

	Comps.Position.Iterate(func(i int, s *Position) {
		if a := Comps.Image.Get(i); a != nil {
			a.op.GeoM.Reset()
			a.op.GeoM.Translate(float64(s.X), float64(s.Y))
			screen.DrawImage(a.image, &a.op)
		}
	})
}
