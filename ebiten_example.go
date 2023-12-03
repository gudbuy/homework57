package main

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"log"
	"image/color"
)

const SCREENHEIGHT, SCREENWIDTH int = 400, 400


type Point struct {
	x, y float64
}

type Vector struct {
	x, y float64
}

func (p Point) Add(v Vector) Point {
	return Point{p.x + v.x, p.y + v.y}
}


type Square struct {
	image *ebiten.Image
	length int
	coords Point
	velocity Vector
}

func CreateSquare(coords Point, length int, velocity Vector) Square {
	square := Square{
		image: ebiten.NewImage(length, length),
		length: length,
		coords: coords,
		velocity: velocity,
	}
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			square.image.Set(x, y, color.White)
		}
	}
	return square
}


func (s *Square) Update() {
	s.coords = s.coords.Add(s.velocity)
}

type Game struct {
	square Square
	op ebiten.DrawImageOptions
}

func (g *Game) Update() error {
	g.square.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(g.square.coords.x, g.square.coords.y)
	screen.DrawImage(g.square.image, &g.op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return SCREENWIDTH, SCREENHEIGHT
}

func main() {
	game := &Game{}
	game.square = CreateSquare(Point{100, 100}, 50, Vector{1, 0})
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Ebiten Example")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}