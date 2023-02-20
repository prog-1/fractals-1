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
	screenWidth  = 640
	screenHeight = 480
)

type Point struct {
	x, y float64
}

type Game struct {
	pos1, pos2    Point
	width, height int
}

func DrawSnowflake(img *ebiten.Image, a, b Point, col color.Color) {
	if math.Abs(b.x-a.x) < 3 && math.Abs(b.y-a.y) < 3 {
		ebitenutil.DrawLine(img, a.x, a.y, b.x, b.y, col)
		return
	}
	x, y := (b.x-a.x)/3, (b.y-a.y)/3
	c := Point{x: a.x + x, y: a.y + y}
	d := Point{x: x*2 + a.x, y: a.y + y*2}
	e := Point{x: (c.x-d.x)*math.Cos(math.Pi/3) - (c.y-d.y)*math.Sin(math.Pi/3) + d.x,
		y: (c.x-d.x)*math.Sin(math.Pi/3) + (c.y-d.y)*math.Cos(math.Pi/3) + d.y}
	DrawSnowflake(img, a, c, col)
	DrawSnowflake(img, c, e, col)
	DrawSnowflake(img, e, d, col)
	DrawSnowflake(img, d, b, col)
}

func (g *Game) Draw(screen *ebiten.Image) {
	e := Point{x: (g.pos1.x-g.pos2.x)*math.Cos(math.Pi/3) - (g.pos1.y-g.pos2.y)*math.Sin(math.Pi/3) + g.pos2.x,
		y: (g.pos1.x-g.pos2.x)*math.Sin(math.Pi/3) + (g.pos1.y-g.pos2.y)*math.Cos(math.Pi/3) + g.pos2.y}
	DrawSnowflake(screen, g.pos2, g.pos1, color.RGBA{R: 227, G: 121, B: 235, A: 255})
	DrawSnowflake(screen, g.pos1, e, color.RGBA{R: 180, G: 255, B: 255, A: 255})
	DrawSnowflake(screen, e, g.pos2, color.RGBA{R: 255, G: 255, B: 140, A: 255})
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		pos1:   Point{100, 200},
		pos2:   Point{300, 200},
	}
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
