package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"runtime"
	"sync"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 500
	viewRadius                = 13
	adjRate                   = 0.015
)

var (
	green = color.RGBA{R: 10, G: 255, B: 50, A: 255}
	boids [boidCount]*Boid
	// represents the space available for the boids. +1 is needed because we count zero.
	boidMap [screenWidth + 1][screenHeight + 1]int
	rWlock  = sync.RWMutex{}
)

type Game struct{}

// Update updates the game
func (g *Game) Update() error {
	return nil
}

// Draw renders the game
func (g *Game) Draw(screen *ebiten.Image) {
	// plot 4 pixels for each boid
	for _, boid := range boids {
		// creates a diamond shape for each boid
		screen.Set(int(boid.position.x+1), int(boid.position.y), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y-1), green)
		screen.Set(int(boid.position.x), int(boid.position.y+1), green)
	}
}

// Layout creates the game area
func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func make2DIntArray() {
	// iterate over rows
	for i, row := range boidMap {
		// populating each element in the row with -1
		for j := range row {
			boidMap[i][j] = -1
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	make2DIntArray()

	// generate Boids
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
