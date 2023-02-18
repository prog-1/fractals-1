package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type (
	point struct {
		x, y float64
	}
	Game struct {
		img *ebiten.Image
		// p1, p2, p3, p4, p5, p6 point
	}
)

const (
	winTitle            = "vectors-4"
	winWidth, winHeight = 500, 500
)

var c = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

func generate(cnt int, img *ebiten.Image, p1, p2 point, sin, cos float64) *ebiten.Image {
	if cnt == 4 {
		ebitenutil.DrawLine(img, p1.x, p1.y, p2.x, p2.y, c)
		return img
	}
	cnt++
	fmt.Println(p1, p2)
	np1, np2 := point{p1.x + (p2.x-p1.x)/3, p1.y + (p2.y-p1.y)/3}, point{p2.x - (p2.x-p1.x)/3, p2.y - (p2.y-p1.y)/3}
	img = generate(cnt, img, p1, np1, sin, cos)
	img = generate(cnt, img, np2, p2, sin, cos)
	ebitenutil.DrawLine(img, p1.x, p1.y, np1.x, np1.y, c)
	ebitenutil.DrawLine(img, np2.x, np2.y, p2.x, p2.y, c)
	newP2x := ((np2.x-np1.x)*cos - (np2.y-np1.y)*sin) + np1.x
	newP2y := ((np2.x-np1.x)*sin + (np2.y-np1.y)*cos) + np1.y
	img = generate(cnt, img, np1, point{newP2x, newP2y}, sin, cos)
	newP1x := ((np2.x-np1.x)*cos - (np2.y-np1.y)*sin) + np1.x
	newP1y := ((np2.x-np1.x)*sin + (np2.y-np1.y)*cos) + np1.y
	img = generate(cnt, img, point{newP1x, newP1y}, np2, sin, cos)
	return img
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.img, nil)
}

func (g *Game) Update() error {
	return nil

}

func (g *Game) Layout(int, int) (w, h int) { return winWidth, winHeight }

func main() {
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowResizable(true)
	img := ebiten.NewImage(winWidth, winHeight)
	p1, p2 := point{100, 400}, point{400, 400}
	var cos = math.Cos(1.04719755)
	var sin = math.Sin(1.04719755)
	img = generate(0, img, p1, p2, sin, cos)
	x, y := p2.x, p2.y
	sin = -sin
	p2.x = (x-p1.x)*cos - (y-p1.y)*sin + p1.x
	p2.y = (x-p1.x)*sin + (y-p1.y)*cos + p1.y
	img = generate(0, img, p1, p2, sin, cos)
	x, y = p1.x, p1.y
	p1.x = (x-p2.x)*cos - (y-p2.y)*sin + p2.x
	p1.y = (x-p2.x)*sin + (y-p2.y)*cos + p2.y
	img = generate(0, img, p1, p2, -sin, cos)
	// img = generate(0, img, point{100, 400}, point{400, 400})
	if err := ebiten.RunGame(&Game{img}); err != nil {
		log.Fatal(err)
	}
}
