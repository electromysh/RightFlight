package main

import (
	"github.com/holtzdari/graphs-app/web"
	_ "github.com/holtzdari/graphs-app/web"
	"log"
	"net/http"
)

func main() {
	web.InitRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
