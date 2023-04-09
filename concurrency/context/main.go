package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type key int

const (
	userIDctx key = 0
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User-Id")

	ctx := context.WithValue(r.Context(), userIDctx, id)
	result := processLonhTask(ctx)
	w.Write([]byte(result))
}

func processLonhTask(ctx context.Context) string {
	id := ctx.Value(userIDctx)

	select {
	case <-time.After(2 * time.Second):
		return fmt.Sprintf("done processing id %d\n", id)
	case <-ctx.Done():
		log.Println("request canceled")
		return ""
	}
}
