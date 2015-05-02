package my

import (
	"net"
)

func Listen(proto, addr string) (chan net.Conn, error) {
	cc := make(chan net.Conn, 10)
	ln, err := net.Listen(proto, addr)
	if err != nil {
		close(cc)
		return cc, err
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				close(cc) //Stop once we hit an error
				return
			}
			cc <- conn
		}
	}()
	return cc, nil
}
