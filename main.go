package main

import (
	"github.com/reflect/ident/idgen"
	"github.com/reflect/ident/source/dynamodb"

	"flag"
	"net/http"
)

func main() {
	var (
		clusterName = flag.String("cluster-name", "development", "name of the cluster")
		tableName   = flag.String("table-name", "Ident", "name of the ident service")
	)

	flag.Parse()

	source := dynamodb.New(*clusterName, *tableName)
	provider, err := idgen.NewProvider(source)
	if err != nil {
		panic(err)
	}

	logInfo("starting ident v%s (id: %d)", Version, provider.Id)

	server := http.Server{
		Addr:    ":8081",
		Handler: NewHandler(provider),
	}

	server.ListenAndServe()
}
