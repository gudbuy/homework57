package main
import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
	ebitenutil "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"image/color"
	"math"
	"fmt"
	)


var SCREENWIDTH, SCREENHEIGHT int = 400, 400
var GRAVITY_ACCELERATION Vector = Vector{0, 0.4}
var GROUND_VELOCITY_LOSS_FACTOR float64 = 0.9


type Point struct {
	x, y float64
}

func (p Point) DistanceTo(other Point) float64 {
	return math.Sqrt((p.x - other.x) * (p.x - other.x) + (p.y - other.y) * (p.y - other.y))
}

func (p Point) AddVector(v Vector) Point {
	return Point{p.x + v.x, p.y + v.y}
}

type Vector struct {
	x, y float64
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func (v Vector) AddVector(other Vector) Vector {
	return Vector{v.x + other.x, v.y + other.y}
}

func (v Vector) Multiply(k float64) Vector {
	return Vector{v.x * k, v.y * k}
}

type Ball struct {
	image *ebiten.Image
	radius int
	coords Point	// center of ball
	velocity, acceleration Vector
}

func CreateBall(coords Point, radius int) Ball {
	ball := Ball{
		image:ebiten.NewImage(radius * 2, radius * 2), 
		radius:radius,
		coords:coords, 
		acceleration:GRAVITY_ACCELERATION,
		velocity:Vector{10, 0},
	}
	
	center := Point{float64(radius), float64(radius)}
	for x := 0; x < radius * 2; x++ {
		for y := 0; y < radius * 2; y++ {
			if center.DistanceTo(Point{float64(x), float64(y)}) <= float64(radius) - 0.1 {
				ball.image.Set(x, y, color.White)
			}
		}			
	}
	
	return ball
}

func (b Ball) IsOnGround() bool {
	return int(b.coords.y) + b.radius == SCREENHEIGHT
}

func (b Ball) IsOnWall() bool {
	return int(b.coords.x) + b.radius == SCREENWIDTH || int(b.coords.x) - b.radius == 0
}

func (b *Ball) Update() {
	if b.IsOnGround() {
		//b.coords = Point{b.coords.x, float64(SCREENHEIGHT - 2 * b.radius)}
		b.velocity = Vector{b.velocity.x, -b.velocity.y}.Multiply(GROUND_VELOCITY_LOSS_FACTOR)
	} else {
		b.velocity = b.velocity.AddVector(b.acceleration)
	}
	
	if b.IsOnWall() {
		b.velocity = Vector{-b.velocity.x, b.velocity.y}
	}
	
	b.coords = b.coords.AddVector(b.velocity)
	
	if b.coords.x < 0 {
		b.coords.x = float64(b.radius)
	}
	if b.coords.x > float64(SCREENWIDTH - b.radius) {
		b.coords.x = float64(SCREENWIDTH - b.radius)
	}
	if b.coords.y > float64(SCREENWIDTH - b.radius) {
		b.coords.y = float64(SCREENHEIGHT - b.radius)
	}
	
}


type Game struct {
	ball Ball
	op ebiten.DrawImageOptions
}

func (g *Game) Update() error {
    g.ball.Update()
    return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
	g.op.GeoM.Reset()
	g.op.GeoM.Translate(g.ball.coords.x - float64(g.ball.radius), g.ball.coords.y - float64(g.ball.radius))
    screen.DrawImage(g.ball.image, &g.op)
    ebitenutil.DebugPrint(screen, fmt.Sprintf("x: %0.2f\ny: %0.2f\nvel: %0.2f\nTPS: %0.2f\nFPS: %0.2f", g.ball.coords.x, g.ball.coords.y, g.ball.velocity.y, ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return SCREENWIDTH, SCREENHEIGHT
}

func main() {
    var game *Game
    game = new(Game)
    game.ball = CreateBall(Point{100, 20}, 15)
    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(800, 800)
    ebiten.SetWindowTitle("Ball")
    //ebiten.SetMaxTPS(120)
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}