package main

import (
    "fmt"
    "math"
    "os"
)

// ToyRobot : state and behaviour container 
type ToyRobot struct {
    position Position 
    bearing Bearing 
}

func (r *ToyRobot) turnLeft() {
    (*r).bearing.rotate(-90)
}

func (r *ToyRobot) turnRight() {
    (*r).bearing.rotate(90)
}

func (r *ToyRobot) move() {
    bearingInRadians := (*r).bearing.toRadians()
    nextPosition := Position{
        (*r).position.x + math.Round(math.Sin(bearingInRadians)),
        (*r).position.y + math.Round(math.Cos(bearingInRadians)),
    }
    if nextPosition.isValid() {
        (*r).position = nextPosition
    }
}

func (r ToyRobot) report() string {
    return fmt.Sprintf("%s,%s", r.position, r.bearing )
}

func (r *ToyRobot) place(p Position, b Bearing) bool {
    if p.isValid() {
        (*r).position = p
        (*r).bearing = b
        return true
    }
    fmt.Fprintln(os.Stderr, "Can not place robot due to invalid position")
    return false
}