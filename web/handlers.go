package web

import (
	"encoding/json"
	"fmt"
	graph "github.com/holtzdari/graphs-app/models"
	"github.com/holtzdari/graphs-app/pkg"
	_ "github.com/holtzdari/graphs-app/pkg"
	"log"
	"net/http"
)

func parse(r *http.Request, dst interface{}) {
	dec := json.NewDecoder(r.Body)

	err := dec.Decode(&dst)

	if err != nil {
		log.Print(err)
	}
}

// Обработчик главной страницы.
func home(w http.ResponseWriter, r *http.Request) {
	var edges []graph.Edge
	parse(r, &edges)
	edgesCount := len(edges)
	countOfEdges, _ := json.Marshal(edgesCount)
	_, err := fmt.Fprintf(w, string(countOfEdges))
	if err != nil {
		return
	}
}

// Обработчик алгоритма Флойда-Уоршелла.
func floydWarshallHandler(w http.ResponseWriter, r *http.Request) {
	var edges []graph.Edge
	parse(r, &edges)
	graph1 := graph.CreateGraph(edges...)
	distances, _ := pkg.FloydWarshallAlgo(graph1)
	allShortestPath, _ := json.Marshal(distances)
	_, err := fmt.Fprintf(w, string(allShortestPath))
	if err != nil {
		return
	}

}

// Обработчик алгоритма Эйлеровых циклов.
func eulerianCyclesHandler(w http.ResponseWriter, r *http.Request) {
	// EulerianCyclesAlgo(............)
}

// Обработчик алгоритма Форда-Фулкерсона.
func fordFulkersonHandler(w http.ResponseWriter, r *http.Request) {
	// FordFulkersonAlgo(.............)
}

// Обработчик алгоритма Прима.
func primHandler(w http.ResponseWriter, r *http.Request) {
	var edges []graph.Edge
	parse(r, &edges)
	graph1 := graph.CreateGraph(edges...)
	mstEdges := pkg.PrimMST(graph1)
	mst, _ := json.Marshal(mstEdges)
	_, err := fmt.Fprintf(w, string(mst))
	if err != nil {
		return
	}
}

func InitRoutes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/floyd-warshall", floydWarshallHandler)
	http.HandleFunc("/eulerian-cycles", eulerianCyclesHandler)
	http.HandleFunc("/ford-fulkerson", fordFulkersonHandler)
	http.HandleFunc("/prim", primHandler)
}
