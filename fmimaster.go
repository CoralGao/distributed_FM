

//
// Task ventilator
// Binds PUSH socket to tcp://localhost:5557
// Sends batch of tasks to workers via that socket
//
package main

import (
    "fmt"
    "github.com/CoralGao/DistSys"

)

type fmimaster struct {
	flag int
}

func (I fmimaster) AnalyzeResult(pattern []byte) {
    fmt.Println("Sync received: ",string(pattern))
}

func main() {
	x := fmimaster{0}
    DistSys.Startmaster(x)
}