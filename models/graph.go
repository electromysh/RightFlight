package models

import (
	"fmt"
)

type Edge struct {
	K string // K for key
	N string // N for neighbour
	W int    // W for weight
}

type Node struct {
	N map[string]int // N for neighbours
	E bool           // E for explored
}

type Graph struct {
	M map[string]Node // M for map
}

func CreateGraph(e ...Edge) Graph {
	g := Graph{M: map[string]Node{}}
	g.AddEdges(e...)
	return g
}

func (g *Graph) AddEdges(edges ...Edge) {
	for _, edge := range edges {
		if _, ok := g.M[edge.K]; !ok {
			g.M[edge.K] = Node{N: map[string]int{}}
		}
		if _, ok := g.M[edge.N]; !ok {
			g.M[edge.N] = Node{N: map[string]int{}}
		}
		g.M[edge.K].N[edge.N] = edge.W
	}
}

func (g *Graph) DeleteEdges(edges ...Edge) {
	for _, edge := range edges {
		delete(g.M[edge.K].N, edge.N)
	}
}

func (g *Graph) Show() {
	fmt.Println(*g)
}

func (g *Graph) GetWeight(key, neighbour string) int {
	if _, ok := g.M[key]; ok {
		if _, ok := g.M[key].N[neighbour]; ok {
			return g.M[key].N[neighbour]
		}
	}
	return 0
}

func (g *Graph) GetEdges(key string) []Edge {
	var edges []Edge
	if _, ok := g.M[key]; ok {
		for neighbour, weight := range g.M[key].N {
			edges = append(edges, Edge{key, neighbour, weight})
		}
	}
	return edges
}

func (g *Graph) GetNeighbours(key string) []string {
	var neighbours []string
	for item := range g.M[key].N {
		neighbours = append(neighbours, item)
	}
	return neighbours
}

func (g *Graph) GetAllEdges() []Edge {
	var edges []Edge
	for key := range g.M {
		for neighbour, weight := range g.M[key].N {
			edges = append(edges, Edge{key, neighbour, weight})
		}
	}
	return edges
}

func (g *Graph) GetAllVertices() []string {
	var vertices []string
	for key := range g.M {
		vertices = append(vertices, key)
	}
	return vertices
}

func (g *Graph) Copy() Graph {
	deepCopy := Graph{map[string]Node{}}
	for key := range g.M {
		deepCopy.M[key] = Node{N: map[string]int{}, E: g.M[key].E}
		for neighbour, weight := range g.M[key].N {
			deepCopy.M[key].N[neighbour] = weight
		}
	}
	return deepCopy
}

func (g *Graph) GetGraph() Graph {
	return *g
}

func (g *Graph) SetE(key string, explored bool) {
	if temp, ok := g.M[key]; ok {
		temp.E = explored
		g.M[key] = temp
	}
}

func (g *Graph) MakeUnvisited() {
	for key := range g.M {
		g.SetE(key, false)
	}
}
