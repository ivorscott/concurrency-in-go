package main

import (
	"math/rand"
	"time"
)

// Boid represents a visual Boid in motion
type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) moveOne() {
	// move boid according to its velocity
	b.position = b.position.Add(b.velocity)
	// calculate next position of the boid
	next := b.position.Add(b.velocity)
	if next.x >= screenWidth || next.x < 0 {
		b.velocity = Vector2D{x: -b.velocity.x, y: b.velocity.y}
	}
	if next.y >= screenHeight || next.y < 0 {
		b.velocity = Vector2D{x: b.velocity.x, y: -b.velocity.y}
	}
}

func (b *Boid) start() {
	// continuously move forever
	// add small delay to slow simulation and give time for other threads to execute
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	b := Boid{
		// pick some random position between 0 and screenWidth or 0 and screenHeight
		position: Vector2D{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
		// limit movement to 1 pixel (-1 or 1), so boids move smoothly without jumping on each update cycle
		velocity: Vector2D{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64() * 2) - 1.0},
		id:       bid,
	}
	boids[bid] = &b
	// start boid movement as a thread
	go b.start()
}
