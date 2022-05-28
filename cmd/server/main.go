package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	voyager "github.com/kiran94/graphql-voyager-go/pkg"
)

func main() {
	const endpoint = "/graphql"
	const port = ":8080"

	schema, err := GenerateSchema()
	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{Schema: schema, GraphiQL: true, Pretty: true})
	vh := voyager.NewVoyagerHandler(endpoint)

	http.Handle(endpoint, h)
	http.Handle("/voyager", vh)

	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
