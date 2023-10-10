package main

import "fmt"

type optFunc func(*opts)

func defaultOpts() *opts {
	return &opts{
		maxConn: 10,
		timeout: 30,
		id:      "1234567890",
		tls:     false,
	}
}

func withTLS(opts *opts) {
	opts.tls = true
}

type opts struct {
	maxConn int
	timeout int
	id      string
	tls     bool
}

type server struct {
	opts opts
}

func newServer(opts ...optFunc) *server {
	o := defaultOpts()
	for _, opt := range opts {
		opt(o)
	}
	return &server{
		opts: *o,
	}
}

func main() {
	s := newServer(withTLS)
	fmt.Println(s)
}
