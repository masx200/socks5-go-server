package socks5_go_server

import "net"

type OnConnectedHandle func(network, address string, port int)
type OnStartedHandle func(conn *net.TCPListener)
