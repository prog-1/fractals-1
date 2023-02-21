package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 700
	screenHeight = 700
)

type Point struct {
	x, y float64
}

type Game struct {
	img           *ebiten.Image
	width, height int
	a, b, e       Point
}

var col = color.RGBA{244, 212, 124, 255}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		a:      Point{100, 500},
		b:      Point{600, 500},
		e:      Point{350, 100},
		img:    ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func fractal(screen *ebiten.Image, a, b Point) {
	if math.Abs(b.x-a.x) < 2 && math.Abs(b.y-a.y) < 2 {
		ebitenutil.DrawLine(screen, a.x, a.y, b.x, b.y, col)
		return
	}
	c := Point{(b.x-a.x)/3 + a.x, (b.y-a.y)/3 + a.y}
	d := Point{(b.x-a.x)/3*2 + a.x, (b.y-a.y)/3*2 + a.y}
	e := Point{(c.x-d.x)*math.Cos(math.Pi/3) - (c.y-d.y)*math.Sin(math.Pi/3) + d.x, (c.x-d.x)*math.Sin(math.Pi/3) + (c.y-d.y)*math.Cos(math.Pi/3) + d.y}

	fractal(screen, a, c)
	fractal(screen, d, b)
	fractal(screen, c, e)
	fractal(screen, e, d)

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	fractal(screen, g.b, g.a)
	fractal(screen, g.e, g.b)
	fractal(screen, g.a, g.e)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
