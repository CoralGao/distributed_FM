

//
// Task Wroker
// Connects PULL socket to tcp://localhost:5557
// Collects workloads from ventilator via that socket
// Connects PUSH socket to tcp://localhost:5558
// Sends results to sink via that socket
//
package main

import ( 
	fmi "github.com/CoralGao/fmindex"
	"github.com/CoralGao/DistSys"
	"strings"
	"fmt"
	"strconv"
)

type fmiworker struct {
	flag int
}

var idx *fmi.Index

func (I fmiworker) Analyze(message []byte) []byte {
	element := strings.Split(string(message), " ")
	if len(element) > 2 {
		fmt.Println("Download the index file.")
		idx = fmi.Load(element[2])
	}
	result := fmi.Search(idx, []byte(element[1]))
	fmt.Println(element[0] + " " + int_string(result))
	return []byte((element[0] + " " + int_string(result)))
}

func int_string(intarray []int) string{
	str := ""
	for i := 0; i < len(intarray); i++ {
		str = str + strconv.Itoa(intarray[i])
	}
	return str
}

func main() {
	x := fmiworker{0}
	DistSys.Startworkers(x)
}
