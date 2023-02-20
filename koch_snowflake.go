package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//---------------------------Declaration--------------------------------

const (
	sW = 640
	sH = 480
)

type Game struct {
	width, height int //screen width and height
	//global variables
	lines []line
}

type point struct {
	x, y float64
}

type line struct {
	start, end point
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	//all logic on update
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	//drawing all the lines in line slice
	for i := range g.lines {
		ebitenutil.DrawLine(screen, g.lines[i].start.x, g.lines[i].start.y, g.lines[i].end.x, g.lines[i].end.y, color.RGBA{255, 255, 255, 255})
	}
}

//-------------------------Functions----------------------------------

//makes koch's snowflake from triangle (returns slice of lines)
func snowflake(a, b, c point, end int) (lines []line) {
	lines = subdivide(a, b, 0, end)
	lines = append(lines, subdivide(b, c, 0, end)...)
	lines = append(lines, subdivide(c, a, 0, end)...)
	return lines
}

//divides line in spiky way _/\_ (returns array of lines)
func subdivide(a, e point, i, end int) (lines []line) {

	//recursion exit condition
	if i > end {
		return
	}
	i++

	//length
	dx := e.x - a.x
	dy := e.y - a.y

	//points
	b := point{a.x + dx/3, a.y + dy/3}
	d := point{e.x - dx/3, e.y - dy/3}
	c := rotate(b, d, -math.Pi/3)

	//recursion for each line segment
	lines = append(lines, subdivide(a, b, i, end)...)
	lines = append(lines, subdivide(b, c, i, end)...)
	lines = append(lines, subdivide(c, d, i, end)...)
	lines = append(lines, subdivide(d, e, i, end)...)

	if i > end { //if it's not pre last time
		//adding lines to line struct
		lines = append(lines, line{a, b})
		lines = append(lines, line{b, c})
		lines = append(lines, line{c, d})
		lines = append(lines, line{d, e})
	}
	return lines

}

//function to rotate the point around another point
//returns new coordinates of the point
func rotate(rp, p point, angle float64) point {
	//p - point, rp - rotation point

	//moving point to top left corner
	p.x, p.y = p.x-rp.x, p.y-rp.y

	//rotating point
	newx := p.x*math.Cos(angle) - p.y*math.Sin(angle)
	newy := p.x*math.Sin(angle) + p.y*math.Cos(angle)

	//returning point that is moved on it's place
	return point{newx + rp.x, newy + rp.y}
}

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {

	//Window
	ebiten.SetWindowSize(sW, sH)
	ebiten.SetWindowTitle("Koch Snowflake")
	ebiten.SetWindowResizable(true) //enablening window resize

	//Game instance
	g := NewGame(sW, sH)                      //creating game instance
	if err := ebiten.RunGame(g); err != nil { //running game
		log.Fatal(err)
	}
}

//New game instance function
func NewGame(width, height int) *Game {

	var a, b, c point
	a.x, a.y = 150, 150
	b.x, b.y = sW-150, 150
	c.x, c.y = sW/2, sH-50

	lines := snowflake(a, b, c, 3)

	// //for triangle debug
	// var lines []line
	// lines = append(lines, line{a, b})
	// lines = append(lines, line{b, c})
	// lines = append(lines, line{c, a})

	return &Game{width: width, height: height, lines: lines}
}
