package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	Height = 480
	Width  = 640
)

// func split(l Line) []Line {

// }

type Game struct {
}

func split(screen *ebiten.Image, x1, y1, x2, y2 float64, depth int) {
	if depth > 7 {
		ebitenutil.DrawLine(screen, x1, y1, x2, y2, color.RGBA{255, 0, 0, 255})
		return
	}
	first_x := (x2 - x1) / 3
	first_y := (y2 - y1) / 3
	split(screen, x1, y1, first_x+x1, first_y+y1, depth+1)
	x := (first_x)*math.Cos(math.Pi/3) - (first_y)*math.Sin(math.Pi/3)
	y := (first_x)*math.Sin(math.Pi/3) + (first_y)*math.Cos(math.Pi/3)
	split(screen, first_x+x1, first_y+y1, x+x1+first_x, y+y1+first_y, depth+1)
	split(screen, x+x1+first_x, y+y1+first_y, x2-first_x, y2-first_y, depth+1)
	split(screen, x2-first_x, y2-first_y, x2, y2, depth+1)

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	split(screen, 100, 100, 300, 300, 0)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func main() {
	ebiten.SetWindowSize(Width, Height)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
