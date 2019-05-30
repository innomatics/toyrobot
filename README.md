# Building

This program is written in Go (tested with v1.12.5). Install the binary release for your OS from https://golang.org/dl/ and follow installation instructions there to install.

Running this command should produce a Go executable `toyrobot`
```
go build
```

# Usage

Simply run the executable.  The simulator will accept commands from standard input e.g.:
```
./toyrobot
> PLACE 1,1,EAST
> MOVE
> REPORT
> 2,1,EAST 
```

# Tests

Run unit tests:
```
go test
```

Some end to end tests are supplied.
```
cd acceptance_testing
./test.sh
```
