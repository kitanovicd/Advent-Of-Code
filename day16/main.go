package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golang-collections/collections/stack"
)

type Valve struct {
	isOpen   bool
	flowRate int
	valves   []string
}

type Collection struct {
	flowRateSum    int
	minutesLeft    int
	lastValveIndex int
	valve          Valve
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	valveGraph := map[string]Valve{}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var flowrate int
		var fromValve string
		line := fileScanner.Text()

		_, err := fmt.Sscanf(line, "Valve %s has flow rate=%d;", &fromValve, &flowrate)
		if err != nil {
			panic("Cannot parse line")
		}

		startIndex := strings.Index(line, "valve") + 7
		endIndex := len(line)
		toValvesConnected := line[startIndex:endIndex]
		toValves := strings.Split(toValvesConnected, ", ")

		valveGraph[fromValve] = Valve{false, flowrate, toValves}
	}

	max := 0
	collection := stack.New()
	collection.Push(Collection{
		flowRateSum:    0,
		minutesLeft:    3,
		lastValveIndex: -1,
		valve:          valveGraph["AA"],
	})

	for collection.Len() > 0 {
		//fmt.Println("Uso")
		currentCollection := collection.Pop().(Collection)
		fmt.Println(currentCollection)
		fmt.Println("Skalnjam ", currentCollection)

		if currentCollection.minutesLeft <= 0 {
			if currentCollection.flowRateSum > max {
				max = currentCollection.flowRateSum
			}
			continue
		}

		nextToVisitIndex := currentCollection.lastValveIndex + 1
		currentCollection.lastValveIndex = nextToVisitIndex
		if nextToVisitIndex < len(currentCollection.valve.valves) {
			var minutesDecrease, flowRateIncrease int
			if currentCollection.valve.isOpen {
				minutesDecrease = 1
				flowRateIncrease = 0
			} else {
				minutesDecrease = 2
				currentCollection.valve.isOpen = true
				flowRateIncrease = currentCollection.valve.flowRate * currentCollection.minutesLeft
			}

			fmt.Println("Dodajem ", currentCollection)
			fmt.Println("Dodajem ", Collection{
				flowRateSum:    currentCollection.flowRateSum + flowRateIncrease,
				minutesLeft:    currentCollection.minutesLeft - minutesDecrease,
				lastValveIndex: -1,
				valve:          valveGraph[currentCollection.valve.valves[nextToVisitIndex]],
			})

			input := bufio.NewScanner(os.Stdin)
			input.Scan()

			collection.Push(currentCollection)
			collection.Push(Collection{
				flowRateSum:    currentCollection.flowRateSum + flowRateIncrease,
				minutesLeft:    currentCollection.minutesLeft - minutesDecrease,
				lastValveIndex: -1,
				valve:          valveGraph[currentCollection.valve.valves[nextToVisitIndex]],
			})
		}
	}

	fmt.Println(max)
}
