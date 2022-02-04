package main

import (
	"math"
	"math/rand"
	"time"
)

// Boid represents a visual Boid in motion
type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)
	// find the enclosing view area using the view radius
	avgPosition, avgVelocity, separation := Vector2D{0, 0}, Vector2D{0, 0}, Vector2D{0, 0}
	count := 0.0

	rWlock.Lock()
	// iterate over the view area with cropping if the view area would otherwise exceed the screen size
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].position)
					separation = separation.Add(b.position.Subtract(boids[otherBoidId].position).DivideV(dist))
				}
			}
		}
	}
	rWlock.Unlock()
	// the rest require no synchronization because we are not reading from other boids
	accel := Vector2D{b.borderBounce(b.position.x, screenWidth), b.borderBounce(b.position.y, screenHeight)}
	if count > 0 {
		avgPosition, avgVelocity = avgPosition.DivideV(count), avgVelocity.DivideV(count)
		accelAlignment := avgVelocity.Subtract(b.velocity).MultiplyV(adjRate)
		accelCohesion := avgPosition.Subtract(b.position).MultiplyV(adjRate)
		accelSeparation := separation.MultiplyV(adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}
	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()
	rWlock.RLock()
	// calculate the alignment of the neighboring boids
	b.velocity = b.velocity.Add(acceleration).limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	// move boid according to its velocity
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	rWlock.RUnlock()
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
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	// start boid movement as a thread
	go b.start()
}
