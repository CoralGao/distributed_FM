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
	"os"
	// "log"
)

type fmiworker struct {
	flag int
}

var idx *fmi.Index
var filename string

func (I fmiworker) Analyze(message []byte) []byte {
	element := strings.Split(string(message), " ")
	if element[0] != filename {
		sequencename := strings.TrimRight(element[0], ".fm")
		if _, err := os.Stat(element[0]); os.IsNotExist(err) {
			if _, err := os.Stat(sequencename); os.IsNotExist(err) {
				return []byte(fmt.Sprintf("No such index: %s, need to build it first. But no such sequence: %s!", filename, sequencename))
				// log.Fatal("Exit!")
			} else {
				filename = element[0]
				fmt.Printf("No such index: %s, build it now!\n", filename)
				idx = fmi.Build(sequencename)
			}
		} else {
			filename = element[0]
			idx = fmi.Load(element[0])
			fmt.Println("Download the index file.")
		}
	}
	result := fmi.Search(idx, []byte(element[2]))
	return format_result(element[1], result)
}

func format_result(index string, intarray []int) []byte{
	str := index
	for i := 0; i < len(intarray); i++ {
		str = str + strconv.Itoa(intarray[i]) + " "
	}
	return []byte(str)
}

func main() {
	x := fmiworker{0}
	DistSys.Startworkers(x)
}
