

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
	"flag"
	"github.com/CoralGao/DistSys"
)

type MyTest struct {
	idx *fmi.Index
}

func (I MyTest) Analyze(pattern []byte) []int {
	return fmi.Search(I.idx, pattern)
}

func main() {
	var index_file = flag.String("i", "", "index file")
	flag.Parse()
	
	if *index_file!="" {
		x := MyTest{}
		x.idx = fmi.Load(*index_file)
		DistSys.Startworkers(x)
	}
}
