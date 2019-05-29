package main

import (
    "testing"
)

func TestToyRobot_turnLeft(t *testing.T) {
    type fields struct {
        position Position
        bearing  Bearing
    }

    tests := []struct {
        name   string
        fields fields
        exp Bearing
    }{
        {
            "Test turn left from North",
            fields{Position{1,1}, Bearing(north)},
            Bearing(west),
        },
        {
            "Test turn left from South",
            fields{Position{1,1}, Bearing(south)},
            Bearing(east),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            r := &ToyRobot{
                position: tt.fields.position,
                bearing:  tt.fields.bearing,
            }
            r.turnLeft()
            if r.bearing != tt.exp {
                t.Errorf("TestToyRobot_turnLeft() %v = %v, want %v",tt.name, r.bearing, tt.exp)
            }
        })
    }
}

func TestToyRobot_turnRight(t *testing.T) {
    type fields struct {
        position Position
        bearing  Bearing
    }
    tests := []struct {
        name   string
        fields fields
        exp Bearing
    }{
        {
            "Test turn right from North",
            fields{Position{1,1}, Bearing(north)},
            Bearing(east),
        },
        {
            "Test turn right from South",
            fields{Position{1,1}, Bearing(south)},
            Bearing(west),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            r := &ToyRobot{
                position: tt.fields.position,
                bearing:  tt.fields.bearing,
            }
            r.turnRight()
            if r.bearing != tt.exp {
                t.Errorf("TestToyRobot_turnRight() %v = %v, want %v",tt.name, r.bearing, tt.exp)
            }
        })
    }
}

func TestToyRobot_move(t *testing.T) {
    type fields struct {
        position Position
        bearing  Bearing
    }
    tests := []struct {
        name   string
        fields fields
        exp Position 
    }{
        {
            "Test move North",
            fields{Position{1,1}, Bearing(north)},
            Position{1,2},
        },
        {
            "Test move East",
            fields{Position{1,1}, Bearing(east)},
            Position{2,1},
        },
        {
            "Test move from East edge (no fall off)",
            fields{Position{4,1}, Bearing(east)},
            Position{4,1},
        },
        {
            "Test move from North edge (no fall off)",
            fields{Position{1,4}, Bearing(north)},
            Position{1,4},
        },
        {
            "Test move from South edge (no fall off)",
            fields{Position{1,0}, Bearing(south)},
            Position{1,0},
        },
        {
            "Test move from West edge (no fall off)",
            fields{Position{0,1}, Bearing(west)},
            Position{0,1},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            r := &ToyRobot{
                position: tt.fields.position,
                bearing:  tt.fields.bearing,
            }
            r.move()
            if r.position != tt.exp {
                t.Errorf("TestToyRobot_move() %v = %v, want %v",tt.name, r.position, tt.exp)
            }
        })
    }
}

func TestToyRobot_report(t *testing.T) {
    type fields struct {
        position Position
        bearing  Bearing
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            "Test Report 1",
            fields{Position{0,1}, Bearing(west)},
            "0,1,WEST",
        },
        {
            "Test Report 2",
            fields{Position{4,3}, Bearing(south)},
            "4,3,SOUTH",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            r := ToyRobot{
                position: tt.fields.position,
                bearing:  tt.fields.bearing,
            }
            if got := r.report(); got != tt.want {
                t.Errorf("ToyRobot.report() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestToyRobot_place(t *testing.T) {
    type fields struct {
        position Position
        bearing  Bearing
    }
    tests := []struct {
        name   string
        fields fields
        want   bool
    }{
        {
            "Test valid placement",
            fields{Position{1,1,}, Bearing(west)},
            true,
        },
        {
            "Test invalid placement",
            fields{Position{5,1,}, Bearing(west)},
            false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            r := &ToyRobot{}
            if got := r.place(tt.fields.position, tt.fields.bearing); got != tt.want {
                t.Errorf("ToyRobot.place() = %v, want %v", got, tt.want)
            }
        })
    }
}
