package main

import (
    "fmt"
    "math"
)

// Bearing : (in degrees) represents the direction robot will move
type Bearing float64

const (
    north Bearing = 0
    east  Bearing = 90
    south Bearing = 180
    west  Bearing = 270
)

// Convert to Radians so we can do trigonometry
func (b Bearing) toRadians() float64 {
    return float64(b) * (math.Pi / 180)
}

// String representation
func (b Bearing) String() string {
    switch b {
    case north:
        return "NORTH"
    case east:
        return "EAST"
    case south:
        return "SOUTH"
    case west:
        return "WEST"
    }
    return fmt.Sprintf("%f", b)
}

// Rotate the bearing by the given degrees
// Output a bearing >=0 and <360 regardless of degrees rotated.
func (b *Bearing) rotate(degrees float64) {
    // handle -ve rotations and multiple rotations
    netPositiveDegrees := 360 + math.Mod(degrees, 360)

    *b = Bearing(math.Mod(float64(*b) + netPositiveDegrees, 360))
}
