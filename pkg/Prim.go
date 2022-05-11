package pkg

import (
	graph "github.com/holtzdari/graphs-app/models"
	"math"
	"math/rand"
)

func PrimMST(g graph.Graph) []graph.Edge {
	var mstEdges []graph.Edge // минимальный остов
	var keys []string         // узлы
	for key := range g.M {
		keys = append(keys, key)
	}
	idx := rand.Intn(len(keys))      // выбираем индекс случайной вершины
	s := keys[idx]                   // случайная вершина
	xMap := map[string]bool{s: true} // помеченные вершины
	for len(xMap) < len(g.M) {       // пока
		minCost := math.MaxInt64 // d[i] <- бесконечность
		var minEdge graph.Edge   // p[i] <- nil
		for key := range xMap {  // перебираем помеченные вершины (ищем минимальное ребро для текущих помеченных)
			for neighbour := range g.M[key].N { // перебирам соседние вершины
				if _, ok := xMap[neighbour]; !ok { // если соседняя вершина НЕ помечена
					if g.M[key].N[neighbour] < minCost { // если ребро для текущей соседней вершины дешевле
						minEdge = graph.Edge{K: key, N: neighbour, W: g.M[key].N[neighbour]} // задаём минимальное ребро
						minCost = g.M[key].N[neighbour]                                      // задаём текущую минимальную стоимость
					}
				}
			}
		}
		xMap[minEdge.N] = true               // помечаем найденный узел
		mstEdges = append(mstEdges, minEdge) // добавляем в остовное дерево найденное ребро
	}
	return mstEdges
}
