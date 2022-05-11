package pkg

import (
	graph "github.com/holtzdari/graphs-app/models"
	"math"
)

func FloydWarshallAlgo(g graph.Graph) ([]graph.Edge, bool) {
	var keys []string
	for key := range g.M {
		keys = append(keys, key)
	}
	cache := make([][]int, len(keys))
	for i := range cache {
		cache[i] = make([]int, len(keys))
		for j := 0; j < len(keys); j++ {
			if i == j {
				cache[i][j] = 0
			} else if val, ok := g.M[keys[i]].N[keys[j]]; ok {
				cache[i][j] = val
			} else {
				cache[i][j] = math.MaxInt32
			}
		}
	}
	var distances []graph.Edge
	for k := 0; k < len(keys); k++ {
		for i := 0; i < len(keys); i++ {
			if i == k {
				continue
			}
			for j := 0; j < len(keys); j++ {
				if j == k {
					continue
				}
				if cache[i][j] > cache[i][k]+cache[k][j] {
					cache[i][j] = cache[i][k] + cache[k][j]
				}
			}
		}
	}
	for i := range keys {
		if cache[i][i] < 0 {
			return distances, false
		}
	}
	for i := range keys {
		for j := range keys {
			if i == j {
				continue
			}
			if float64(cache[i][j])/float64(math.MaxInt32) < 0.5 {
				distances = append(distances, graph.Edge{K: keys[i], N: keys[j], W: cache[i][j]})
			}
		}
	}
	return distances, true
}
