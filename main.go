package main

import (
	"fmt"
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

type Line struct {
	offset_x, offset_y float64
	rotation           float64
	length             float64
}

func (l *Line) Draw(screen *ebiten.Image) {
	// fmt.Println(l.rotation)
	x1, y1 := l.offset_x, l.offset_y
	x2, y2 := (x1+l.length)*math.Cos(l.rotation)-(y1-l.offset_y)*math.Sin(l.rotation)+l.length, (x1)*math.Sin(l.rotation)+(y1-l.offset_y)*math.Cos(l.rotation)+l.offset_y
	ebitenutil.DrawLine(screen, x1, y1, x2, y2, color.RGBA{255, 0, 0, 255})
}

// func split(l Line) []Line {

// }

type Game struct {
	Lines []Line
}

func split(l Line) (Lines [4]Line) {
	Lines[0].offset_x = l.offset_x
	Lines[0].offset_y = l.offset_y
	Lines[0].length = l.length / 4

	Lines[1].offset_x = l.offset_x + (l.length / 4)
	Lines[1].offset_y = l.offset_y
	Lines[1].length = l.length / 4
	Lines[1].rotation = l.rotation - math.Pi/3

	// Lines[2].offset_x = l.offset_x + 2*(l.length/10)
	// Lines[2].offset_y = l.offset_y
	// Lines[2].length = l.length / 10
	// Lines[2].rotation = l.rotation + math.Pi/3
	return
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range split(g.Lines[0]) {
		v.Draw(screen)
		fmt.Print(v.offset_x)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func main() {
	ebiten.SetWindowSize(Width, Height)
	if err := ebiten.RunGame(&Game{Lines: []Line{Line{200, 200, 0, 100}}}); err != nil {
		log.Fatal(err)
	}
}
