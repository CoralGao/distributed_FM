

//
// Task ventilator
// Binds PUSH socket to tcp://localhost:5557
// Sends batch of tasks to workers via that socket
//
package main

import (
    "fmt"
    "github.com/CoralGao/DistSys"
    "strconv"
)

type fmimaster struct {
	flag int
}

func (I fmimaster) AnalyzeResult(pattern []byte) {
    fmt.Println("Sync received: ",string(pattern))
}

func (I fmimaster) ProduceMsg(line []byte, count int, filename string) []byte {
        line = line[0:len(line)-1]
        msg := filename + " " + strconv.Itoa(count+1) + " " + string(line)
        return []byte(msg)
}

func main() {
	x := fmimaster{0}
    DistSys.Startmaster(x)
}
