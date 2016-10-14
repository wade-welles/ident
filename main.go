package main

import (
	"flag"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Run(h *Handler, ch chan string) {
	for {
		ch <- h.generateToken()
	}
}

func main() {
	var (
		clusterName = flag.String("cluster-name", "development", "name of the cluster")
		tableName   = flag.String("table-name", "Ident", "name of the ident service")
	)

	flag.Parse()

	id := MustID(*clusterName, *tableName)

	logInfo("starting ident v%s (id: %d)", Version, id)

	server := http.Server{
		Addr:    ":8081",
		Handler: NewHandler(id),
	}

	server.ListenAndServe()
}
