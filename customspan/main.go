package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/siuyin/dflt"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("custom-span-server")

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), "/ entry")
		defer span.End()
		io.WriteString(w, "/ entry\n")
		subProc(ctx, w, r)
	})

	port := dflt.EnvString("PORT", "8080")
	log.Println("Starting server. PORT=", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func subProc(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	log.Println("subProc started")
	_, span := tracer.Start(ctx, "subProc")
	io.WriteString(w, "subProc\n")
	log.Println(r)
	defer span.End()
	log.Println("subProc completed")
}
