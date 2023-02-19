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
	width, height int
	a, b          Point
}

func Rotate(a Point, b Point, angle float64) Point {
	return Point{x: (a.x-b.x)*math.Cos(angle) - (a.y-b.y)*math.Sin(angle) + b.x, y: (a.x-b.x)*math.Sin(angle) + (a.y-b.y)*math.Cos(angle) + b.y}
}

func DrawKochSnowflake(img *ebiten.Image, a, b Point) {
	if math.Abs(b.x-a.x) < 5 && math.Abs(b.y-a.y) < 5 {
		ebitenutil.DrawLine(img, a.x, a.y, b.x, b.y, color.White)
		return
	}
	x, y := (b.x-a.x)/3, (b.y-a.y)/3
	c := Rotate(Point{x: a.x + x, y: a.y + y}, Point{x: b.x - x, y: b.y - y}, math.Pi/3)
	DrawKochSnowflake(img, a, Point{x: a.x + x, y: a.y + y})
	DrawKochSnowflake(img, Point{x: a.x + x, y: a.y + y}, c)
	DrawKochSnowflake(img, c, Point{x: b.x - x, y: b.y - y})
	DrawKochSnowflake(img, Point{x: b.x - x, y: b.y - y}, b)
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		a:      Point{200, 300},
		b:      Point{500, 400},
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	c := Rotate(g.a, g.b, math.Pi/3)
	DrawKochSnowflake(screen, g.b, g.a)
	DrawKochSnowflake(screen, g.a, c)
	DrawKochSnowflake(screen, c, g.b)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
