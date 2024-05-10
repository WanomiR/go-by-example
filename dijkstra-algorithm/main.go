package main

import (
	"fmt"
	"math"
)

type node map[string]float64

type Grapher interface {
	FindClosestNode(string) string
	WalkGrpah() float64
}

type Graph struct {
	graph     map[string]node
	distances map[string]float64
	parents   map[string]string
	processed map[string]bool
}

func main() {
	g := NewGraph()

	g.graph["start"] = node{"a": 5, "b": 2}
	g.graph["a"] = node{"c": 4, "d": 2}
	g.graph["b"] = node{"a": 8, "d": 7}
	g.graph["c"] = node{"d": 6, "finish": 3}
	g.graph["d"] = node{"finish": 1}

	g.distances["a"] = 5
	g.distances["b"] = 2
	g.distances["c"] = math.Inf(1)
	g.distances["d"] = math.Inf(1)
	g.distances["finish"] = math.Inf(1)

	g.parents["a"] = "start"
	g.parents["b"] = "start"

	fmt.Println(g.WalkGrpah())
}

func (g *Graph) WalkGrpah() float64 {
	currentNode := g.FindClosestNode()

	for currentNode != "" {
		dist := g.distances[currentNode]
		neighbors := g.graph[currentNode]

		for n, d := range neighbors {
			newDist := dist + d

			if newDist < g.distances[n] {
				g.distances[n] = newDist
				g.parents[n] = currentNode
			}
		}

		g.processed[currentNode] = true
		currentNode = g.FindClosestNode()
	}

	return g.distances["finish"]
}

func (g *Graph) FindClosestNode() string {
	var closestNode string
	minDist := math.Inf(1)
	for currentNode, dist := range g.distances {
		// if not in processed nodes and with less distance
		if _, ok := g.processed[currentNode]; !ok && dist < minDist {
			closestNode = currentNode
			minDist = dist
		}
	}
	return closestNode
}

func NewGraph() *Graph {
	return &Graph{
		graph:     make(map[string]node),
		distances: make(map[string]float64),
		parents:   make(map[string]string),
		processed: make(map[string]bool),
	}
}
