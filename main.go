package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/dgravesa/WaterLogger-UserServer/controller"
	"github.com/dgravesa/WaterLogger-UserServer/data"
	"github.com/dgravesa/WaterLogger-UserServer/model"
)

var port = flag.Uint("port", 11011, "port to listen on")

func main() {
	flag.Parse()

	controller.RegisterRoutes()

	log.Println("initializing data layer")
	model.SetUserDataLayer(data.NewInMemoryUserStore())

	log.Printf("listening on port %d", *port)
	portStr := fmt.Sprintf(":%d", *port)
	http.ListenAndServe(portStr, nil)
}
