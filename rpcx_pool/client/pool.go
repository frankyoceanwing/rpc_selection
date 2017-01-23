package main

import (
	"errors"
	"sync"
	"time"

	"fmt"

	"github.com/smallnest/rpcx"
)

// Pool holds clients.
type Pool struct {
	mu             sync.RWMutex
	clients        chan *rpcx.Client
	clientSelector rpcx.ClientSelector
	config         *Config
	size           int
}

// Config is used to configure the pool.
type Config struct {
	InitCap int
	MaxCap  int
	Timeout time.Duration
}

// NewPool creates a new pool of clients.
func NewPool(c *Config, s rpcx.ClientSelector) (*Pool, error) {
	if c == nil {
		return nil, errors.New("config should not be nil")
	}
	if c.InitCap < 0 || c.MaxCap <= 0 || c.InitCap > c.MaxCap {
		return nil, errors.New("invalid capacity")
	}

	p := &Pool{
		clients:        make(chan *rpcx.Client, c.MaxCap),
		clientSelector: s,
		config:         c,
	}
	p.init()
	return p, nil
}

func (p *Pool) init() {
	for i := 0; i < p.config.InitCap; i++ {
		p.clients <- p.newClient()
	}
}

// Borrow a client from the pool.
func (p *Pool) Borrow() *rpcx.Client {
	var c *rpcx.Client
	select {
	case c = <-p.clients:
		return c
	default: // pool is empty
		return p.tryNewClient()
	}
}

func (p *Pool) tryNewClient() *rpcx.Client {
	var c *rpcx.Client
	if p.getSize() >= p.config.MaxCap {
		select {
		case c = <-p.clients:
			return c
		case <-time.After(p.config.Timeout):
			fmt.Printf("timeout\n")
			return nil
		}
	}
	return p.newClient()
}

func (p *Pool) getSize() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.size
}

func (p *Pool) newClient() *rpcx.Client {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.size++
	return rpcx.NewClient(p.clientSelector)
}

// Return returns a client to the pool.
func (p *Pool) Return(c *rpcx.Client) {
	select {
	case p.clients <- c:
	default: // pool is filled up, throw it
	}
}

// Throw throws away a client.
func (p *Pool) Throw(c *rpcx.Client) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.size--
}

// Clear closes all clients and removes them from the pool.
func (p *Pool) Clear() {
	// TODO close and remove clients
}
