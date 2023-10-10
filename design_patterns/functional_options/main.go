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

func withMaxConn(maxConn int) optFunc {
	return func(opts *opts) {
		opts.maxConn = maxConn
	}
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
	s := newServer(withTLS, withMaxConn(20))
	fmt.Println(s)
}
