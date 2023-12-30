// https://www.youtube.com/watch?v=MDy7JQN5MN4

package main

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

func SetMaxConn(i int) OptFunc {
	return func(o *Opts) {
		o.maxConn = i
	}
}

func EnableTLS(o *Opts) {
	o.tls = true
}

type Server struct {
	Opts
}

func NewServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		o,
	}
}

func main() {
	server := NewServer(SetMaxConn(100), EnableTLS)
	fmt.Println(server)
}
