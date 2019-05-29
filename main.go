package main

import (
    "bufio"
    "errors"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

/*
 * Toy Robot simulator - process these commands from standard input.
 * 
 * Commands are recognised as the first word in each line of input (case-sensitive).
 * Mostly other input is ignored, however the program will exit with error if the 
 * PLACE command has invalid arguments or format.
 */
const (
    PLACE  string = "PLACE"  // Place a robot on the table top e.g. PLACE 2,3,WEST 
    LEFT   string = "LEFT"   // Turn the robot 90 degress to it's left
    RIGHT  string = "RIGHT"  // Turn the robot 90 degress to it's right
    MOVE   string = "MOVE"   // Move the robot forward one space
    REPORT string = "REPORT" // Report the robot's current position and bearing e.g 3,3,NORTH
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var robot *ToyRobot
    for scanner.Scan() {
        robot = processLine(scanner.Text(), robot)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

// Process a line of input, detect commands and direct them to a robot
// Testing is covered by acceptance test
func processLine(line string, robot *ToyRobot) *ToyRobot {
    tokens := strings.Split(line, " ")
    if len(tokens) > 0 {
        if tokens[0] == PLACE {
            position, bearing, err := parsePlaceCommand(tokens)
            if err != nil {
                log.Fatal(err)
            }
            
            // Attempt to place a new robot
            newRobot := ToyRobot{}
            if newRobot.place(position, bearing) {
                // A new robot has been placed, so return reference to that.
                // The reference to any older robot will be garbage collected.
                return &newRobot
            }
        } else {
            // Skip other commands if robot has not been placed yet
            if robot != nil {
                switch tokens[0] {
                case LEFT:
                    (*robot).turnLeft()
                case RIGHT:
                    (*robot).turnRight()
                case MOVE:
                    (*robot).move()
                case REPORT:
                    fmt.Fprintln(os.Stdout, (*robot).report())
                }
            }
        }
    }
    return robot
}

// Parse the arguments for the PLACE command
func parsePlaceCommand(cmd []string) (Position, Bearing, error) {
    var p Position
    var b Bearing
    if len(cmd) < 2 {
        return p, b, errors.New("Arguments not supplied with PLACE command")
    }
    args := strings.Split(cmd[1], ",")
    if len(args) != 3 {
        return p, b, errors.New("Wrong number of arguments supplied with PLACE command (expected 3)")
    }
    x, err := strconv.Atoi(args[0])
    if err != nil {
        return p, b, errors.New("Invalid first argument to PLACE command (expected integer)")
    }
    y, err := strconv.Atoi(args[1])
    if err != nil {
        return p, b, errors.New("Invalid second argument to PLACE command (expected integer)")
    }
    switch args[2] {
    case "NORTH":
        b = north
    case "EAST":
        b = east
    case "SOUTH":
        b = south
    case "WEST":
        b = west
    default:
        return p, b, errors.New("Invalid third argument to PLACE command")
    }
    return Position{float64(x), float64(y)}, b, nil
}
