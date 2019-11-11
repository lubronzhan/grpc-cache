package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	rpc "github.com/lubronzhan/grpc-cache/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CacheService struct {
	store         map[string][]byte
	mu            sync.Mutex
	accounts      rpc.AccountsClient
	keysByAccount map[string]int64
}

func NewCacheService() *CacheService {
	return &CacheService{
		store:         make(map[string][]byte),
		keysByAccount: make(map[string]int64),
	}
}

func (s *CacheService) Get(ctx context.Context, req *rpc.GetReq) (*rpc.GetResp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.store[req.Key]
	if !ok {
		return nil, fmt.Errorf("Key '%s' not found", req.Key)
	}

	return &rpc.GetResp{Val: val}, nil
}

func (s *CacheService) Store(ctx context.Context, req *rpc.StoreReq) (*rpc.StoreResp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	start := time.Now()
	resp, err := s.accounts.GetByToken(context.Background(), &rpc.GetByTokenReq{
		Token: req.AccountToken,
	})
	log.Printf("accounts.GetByToken duration %s", time.Since(start))
	if err != nil {
		return nil, err
	}

	if s.keysByAccount[req.AccountToken] >= resp.Account.MaxCacheKeys {
		return nil, status.Errorf(codes.FailedPrecondition, "Account %s exceeds max key limit %d", req.AccountToken, resp.Account.MaxCacheKeys)
	}

	if _, ok := s.store[req.Key]; !ok {
		s.keysByAccount[req.AccountToken]++
	}

	s.store[req.Key] = req.Val

	return &rpc.StoreResp{}, nil
}

func runServer() error {
	srv := grpc.NewServer()
	cacheService := NewCacheService()
	rpc.RegisterCacheServer(srv, cacheService)
	l, err := net.Listen("tcp", "localhost:5051")
	if err != nil {
		return err
	}

	conn, err := grpc.Dial("localhost:5052", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to dial accounts server: %v", err)
	}

	cacheService.accounts = rpc.NewAccountsClient(conn)

	return srv.Serve(l)
}

func main() {
	if err := runServer(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run cache server: %s\n", err)
		os.Exit(1)
	}
}
