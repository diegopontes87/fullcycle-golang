package main

import (
	"context"
	"fmt"
)

// Define a custom type for context keys to avoid collisions
type contextKey string

const tokenKey = contextKey("token")

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, tokenKey, "password")
	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	token := ctx.Value(tokenKey)
	fmt.Println(token)
}
