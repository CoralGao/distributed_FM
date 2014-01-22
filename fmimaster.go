

//
// Task ventilator
// Binds PUSH socket to tcp://localhost:5557
// Sends batch of tasks to workers via that socket
//
package main

import (
        "fmt"
        master "github.com/CoralGao/DistSys/master"

)

type fmimaster struct {
	flag int
}

func (I fmimaster) Analyze(pattern []byte) {
    fmt.Println("Sync received: ",string(pattern))
}

func main() {
	x := fmimaster{0}
    master.Start(x)
}