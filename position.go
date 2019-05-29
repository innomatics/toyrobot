package main

import (
    "fmt"
)

// Position : represents a location on the Table Top
type Position struct {
    x float64
    y float64
}

const tableTopSize = 5

func (p Position) String() string {
    return fmt.Sprintf("%d,%d", int(p.x), int(p.y))
}

func (p Position) isValid() bool {
    return p.x < tableTopSize && p.y < tableTopSize && p.x >= 0 && p.y >= 0
}
