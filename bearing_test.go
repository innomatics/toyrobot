package main

import (
    "testing"
    "math"
)

func TestBearing_rotate(t *testing.T) {
    type args struct {
        degrees float64
    }

    bearingZero := Bearing(0)
    bearingNorth := Bearing(north)
    bearingWest:= Bearing(west)
    bearingWest1:= Bearing(west)

    tests := []struct {
        name string
        b    *Bearing
        args args
        exp Bearing
    }{
        {
            "Test ninety degree rotation from zero",
            &bearingZero,
            args{90},
            Bearing(east),
        },
        {
            "Test minus ninety degree rotation from north",
            &bearingNorth,
            args{-90},
            Bearing(west),
        },
        {
            "Test ninety degree rotation from west",
            &bearingWest,
            args{90},
            Bearing(north),
        },
        {
            "Test minus ninety degree rotation from west",
            &bearingWest1,
            args{-90},
            Bearing(south),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.b.rotate(tt.args.degrees)
        })
        if *(tt.b) != tt.exp {
            t.Errorf("TestBearing_rotate() %v = %v, want %v",tt.name, *(tt.b), tt.exp)
        }
    }
}

func TestBearing_String(t *testing.T) {
    tests := []struct {
        name string
        b    Bearing
        want string
    }{
        {
            "test bearing string N",
            Bearing(north),
            "NORTH",
        },
        {
            "test bearing string S",
            Bearing(south),
            "SOUTH",
        },
        {
            "test bearing string E",
            Bearing(east),
            "EAST",
        },
        {
            "test bearing string W",
            Bearing(west),
            "WEST",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.b.String(); got != tt.want {
                t.Errorf("Bearing.String() %v = %v, want %v", tt.name, got, tt.want)
            }
        })
    }
}

func TestBearing_toRadians(t *testing.T) {
    tests := []struct {
        name string
        b    Bearing
        want float64
    }{
        {
            "Test 90 degress",
            Bearing(90),
            math.Pi/2,
        },
        {
            "Test 180 degress",
            Bearing(180),
            math.Pi,
        },
        {
            "Test 270 degress",
            Bearing(270),
            math.Pi*1.5,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.b.toRadians(); got != tt.want {
                t.Errorf("Bearing.toRadians() = %v, want %v", got, tt.want)
            }
        })
    }
}
