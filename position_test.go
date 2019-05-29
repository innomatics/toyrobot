package main

import (
    "testing"
)

func TestPosition_isValid(t *testing.T) {
    type fields struct {
        x float64
        y float64
    }
    tests := []struct {
        name   string
        fields fields
        want   bool
    }{
        {
            "valid position SE corner",
            fields{0, 0},
            true,
        },
        {
            "valid position NW corner",
            fields{0, 4},
            true,
        },
        {
            "valid position NE corner",
            fields{4, 4},
            true,
        },
        {
            "valid position SW corner",
            fields{4, 0},
            true,
        },
        {
            "invalid position SE corner 1",
            fields{-1, 0},
            false,
        },
        {
            "invalid position SE corner 2",
            fields{0, -1},
            false,
        },
        {
            "invalid position SE corner 3",
            fields{-1, -1},
            false,
        },
        {
            "invalid position NW corner 1",
            fields{0, 5},
            false,
        },
        {
            "invalid position NW corner 2",
            fields{-1, 4},
            false,
        },
        {
            "invalid position NW corner 3",
            fields{-1, 5},
            false,
        },
        {
            "invalid position NE corner 1",
            fields{4, 5},
            false,
        },
        {
            "invalid position NE corner 2",
            fields{5, 4},
            false,
        },
        {
            "invalid position NE corner 3",
            fields{5, 5},
            false,
        },
        {
            "invalid position SW corner 1",
            fields{5, 0},
            false,
        },
        {
            "invalid position SW corner 2",
            fields{4, -1},
            false,
        },
        {
            "invalid position SW corner 3",
            fields{5, -1},
            false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            p := Position{
                x: tt.fields.x,
                y: tt.fields.y,
            }
            if got := p.isValid(); got != tt.want {
                t.Errorf("Position.isValid() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPosition_String(t *testing.T) {
    type fields struct {
        x float64
        y float64
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            "position string representation",
            fields{0, 3},
            "0,3",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            p := Position{
                x: tt.fields.x,
                y: tt.fields.y,
            }
            if got := p.String(); got != tt.want {
                t.Errorf("Position.String() = %v, want %v", got, tt.want)
            }
        })
    }
}
