package main

import (
    "reflect"
    "testing"
)

func Test_parsePlaceCommand(t *testing.T) {
    type args struct {
        cmd []string
    }
    tests := []struct {
        name    string
        args    args
        want    Position
        want1   Bearing
        wantErr bool
    }{
        {
            "valid PLACE command 1",
            args{[]string{"PLACE", "3,4,WEST"}},
            Position{3, 4},
            Bearing(west),
            false,
        },
        {
            "valid PLACE command 2",
            args{[]string{"PLACE", "1,1,EAST"}},
            Position{1, 1},
            Bearing(90),
            false,
        },
        {
            "invalid PLACE command 1",
            args{[]string{"PLACE", "1,1,SOUTHEAST"}},
            Position{0, 0},
            Bearing(0),
            true,
        },
        {
            "invalid PLACE command 2",
            args{[]string{"PLACE", "1 1 EAST"}},
            Position{0, 0},
            Bearing(0),
            true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, got1, err := parsePlaceCommand(tt.args.cmd)
            if (err != nil) != tt.wantErr {
                t.Errorf("parsePlaceCommand() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("parsePlaceCommand() got = %v, want %v", got, tt.want)
            }
            if got1 != tt.want1 {
                t.Errorf("parsePlaceCommand() got1 = %v, want %v", got1, tt.want1)
            }
        })
    }
}
