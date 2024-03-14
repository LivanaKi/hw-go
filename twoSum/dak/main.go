package main

import(
	"fmt"
	"math"
)

var Processed = map[string]bool{}
const endNode = "NONE"

func main() {
	type graph map[string]map[string]int
	var g = graph{
		"start": map[string]int{"a": 6, "b": 2},
		"a":     map[string]int{"fin": 1},
		"b":     map[string]int{"a": 3, "fin": 5},
		"fin":   map[string]int{},
	}

	var costs = map[string]int{
		"a":   6,
		"b":   2,
		"fin": math.MaxInt32,
	}

	var parents = map[string]string{
		"a":   "start",
		"b":   "start",
		"fin": endNode,
	}

	node := findLowestCostNode(costs)

	for node != "" {
		cost := costs[node]
		neighbors := g[node]
		for n, d := range neighbors {
			newCost := cost + d
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		Processed[node] = true
		node = findLowestCostNode(costs)
	}
	fmt.Println(costs)
	fmt.Println(parents)
}

func findLowestCostNode(costs map[string]int) string {
	node := ""
	low := math.MaxInt32
	for k, v := range costs {
		if _, ok := Processed[k]; !ok && v < low {
			low = v
			node = k
		}
	}
	return node
}
