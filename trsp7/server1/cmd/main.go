package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
	"log"
	"net/http"
	"os"
	db2 "practice4/db"
	"practice4/graph"
	handler2 "practice4/handler"
	"sync"
)

const defaultPort = "10001"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()

		db, err := db2.InitDataBase()
		if err != nil {
			log.Fatal(err)
		}

		port := os.Getenv("PORT")
		if port == "" {
			port = defaultPort
		}

		srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{db}}))

		srv.AddTransport(transport.Options{})
		srv.AddTransport(transport.GET{})
		srv.AddTransport(transport.POST{})

		srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

		srv.Use(extension.Introspection{})
		srv.Use(extension.AutomaticPersistedQuery{
			Cache: lru.New[string](100),
		})

		middleware := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://127.0.0.1:3000"},
			AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
			AllowedHeaders:   []string{"Content-Type", "Authorization"},
			AllowCredentials: true,
		})

		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", middleware.Handler(srv))

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()
	go func() {
		handler2.InitRoutes()
		defer wg.Done()
	}()

	wg.Wait()
}
