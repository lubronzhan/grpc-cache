package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	rpc "github.com/lubronzhan/grpc-cache/protobuf"
	"google.golang.org/grpc"
)

func main() {
	if err := runClient(); err != nil {
		fmt.Fprintf(os.Stderr, "failed: %v\n", err)
		os.Exit(1)
	}
}

func runClient() error {
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to dail server: %v", err)
	}

	cache := rpc.NewCacheClient(conn)

	start := time.Now()
	_, err = cache.Store(context.Background(),
		&rpc.StoreReq{
			Key:          "gopher",
			Val:          []byte("con"),
			AccountToken: "inconshreveable",
		})
	log.Printf("cache.Store duration %s", time.Since(start))

	if err != nil {
		return fmt.Errorf("failed to store: %v", err)
	}

	start = time.Now()
	resp, err := cache.Get(context.Background(), &rpc.GetReq{Key: "gopher"})
	log.Printf("cache.Get duration %s", time.Since(start))
	if err != nil {
		return fmt.Errorf("failed to get: %v", err)
	}

	fmt.Printf("Got cached value %s\n", resp.Val)

	return nil
}
