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

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func TestSys() {

	hmm := 5
	what := false

	Positions.Iterate(func(s *Position) {
		s.X = s.X + hmm
		print(s.X)
		if s.X > 500 {
			what = true
		}
	})

	if what {
		Positions.Iterate(func(s *Position) {
			print("hejsan")
		})
	}

}
