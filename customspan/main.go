package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/siuyin/dflt"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("custom-span-server")

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), "/ entry")
		slog.Info("/ entry", "path", r.URL.Path)
		defer span.End()
		io.WriteString(w, "/ entry\n")
		subProc(ctx, w, r)
	})

	port := dflt.EnvString("PORT", "8080")
	log.Println("Starting server. PORT=", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func subProc(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(ctx, "subProc")
	defer span.End()
	slog.Info("subProc", "started", time.Now())
	io.WriteString(w, "subProc\n")
	slog.Info("subProc", "completed", time.Now())
}
