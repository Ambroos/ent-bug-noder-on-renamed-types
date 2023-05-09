package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/migrate"
	"entgo.io/bug/graph"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	client = client.Debug()

	ctx := context.Background()

	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	client.User.Delete().ExecX(ctx)
	u1 := client.User.Create().SetName("Ariel").SetAge(30).SaveX(ctx)
	client.UnrenamedUser.Delete().ExecX(ctx)
	u2 := client.UnrenamedUser.Create().SetName("Ariel").SetAge(30).SaveX(ctx)

	srv := handler.NewDefaultServer(graph.NewSchema(client))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Printf("renamed user ID: %d", u1.ID)
	log.Printf("unrenamed user ID: %d", u2.ID)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
