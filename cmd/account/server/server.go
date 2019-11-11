package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"sync"

	rpc "github.com/lubronzhan/grpc-cache/protobuf"
	"google.golang.org/grpc"
)

type TokenService struct {
	accounts map[string]int64
	mu       sync.Mutex
}

func NewTokenService() *TokenService {
	return &TokenService{
		accounts: make(map[string]int64),
	}
}

func (t *TokenService) GetByToken(ctx context.Context, req *rpc.GetByTokenReq) (*rpc.GetByTokenResp, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, ok := t.accounts[req.Token]; !ok {
		fmt.Printf("Generating random limit for user %s", req.Token)
		t.accounts[req.Token] = rand.Int63n(30)
	}

	return &rpc.GetByTokenResp{
		Account: &rpc.Account{
			MaxCacheKeys: t.accounts[req.Token],
		},
	}, nil
}

func runServer() error {
	srv := grpc.NewServer()
	rpc.RegisterAccountsServer(srv, NewTokenService())
	l, err := net.Listen("tcp", "localhost:5052")
	if err != nil {
		return err
	}

	return srv.Serve(l)
}

func main() {
	if err := runServer(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run token server: %s\n", err)
		os.Exit(1)
	}
}
