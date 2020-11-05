package logger

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

const requestIDKey string = "request-id"

func Println(ctx context.Context, msg string) {
	requestID, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		log.Println("could not find request ID in context")
		return
	}

	log.Printf("[%d] %s\n", requestID, msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)
		f(rw, r.WithContext(ctx))
	}
}
