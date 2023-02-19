package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type vector struct {
	x, y float64
}

func (a *vector) mod() float64 {
	return math.Sqrt(math.Pow(a.x, 2) + math.Pow(a.y, 2))
}

func subtract(a, b *vector) vector {
	return vector{a.x - b.x, a.y - b.y}
}

type line struct {
	a, b vector
}

func (l *line) draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, l.a.x, l.a.y, l.b.x, l.b.y, color.RGBA{0, 255, 255, 255})
}

type game struct {
	l            line
	screenBuffer *ebiten.Image
}

func rotate(l *line, rad float64) vector {
	x, y := l.b.x-l.a.x, l.b.y-l.a.y
	x, y = x*math.Cos(rad)-y*math.Sin(rad), y*math.Cos(rad)+x*math.Sin(rad)
	return vector{x + l.a.x, y + l.a.y}
}

func NewGame() *game {
	return &game{
		line{vector{screenWidth / 3, screenHeight / 3}, vector{screenWidth / 3 * 2, screenHeight / 3}},
		ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.screenBuffer, &ebiten.DrawImageOptions{})
}

func KochSnowflake(l line, screen *ebiten.Image) {
	if dif := subtract(&l.b, &l.a); dif.mod() < 5 {
		l.draw(screen)
		return
	}
	a, b := l.a, l.b
	c := vector{(b.x-a.x)/3 + a.x, (b.y-a.y)/3 + a.y}
	KochSnowflake(line{a, c}, screen)

	d := vector{(b.x-a.x)/3*2 + a.x, (b.y-a.y)/3*2 + a.y}
	KochSnowflake(line{d, b}, screen)

	e := rotate(&line{c, d}, -math.Pi/3)
	KochSnowflake(line{c, e}, screen)

	KochSnowflake(line{e, d}, screen)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame()
	KochSnowflake(g.l, g.screenBuffer)
	KochSnowflake(line{rotate(&g.l, math.Pi/3), g.l.a}, g.screenBuffer)
	KochSnowflake(line{g.l.b, rotate(&g.l, math.Pi/3)}, g.screenBuffer)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
