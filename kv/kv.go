package kv

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	Key string
	Val string
}

type StatusType int

const (
	Success = iota
	Fail
)

type Reply struct {
	Status StatusType
	Val    string
}

type Server struct {
	kv map[string]string
}

func (s *Server) init() {
	s.kv = map[string]string{}
}

func (s *Server) Set(args *Args, reply *Reply) error {
	key, val := args.Key, args.Val
	s.kv[key] = val
	reply.Status = Success
	reply.Val = ""
	return nil
}

func (s *Server) Get(args *Args, reply *Reply) error {
	key := args.Key
	val, ok := s.kv[key]
	if !ok {
		reply.Status = Fail
		reply.Val = ""
		return nil
	} else {
		reply.Status = Success
		reply.Val = val
		return nil
	}
}

func RunServer(addr string) {
	server := new(Server)
	server.init() // 外部不可见
	rpc.Register(server)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Listen Error:", err)
	}
	http.Serve(listener, nil)
}

type Client struct {
	client *rpc.Client
}

func NewClient(addr string) *Client {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	c := &Client{client: client}
	return c
}

func (c *Client) Set(key string, val string) StatusType {
	args := &Args{Key: key, Val: val}
	reply := new(Reply)
	err := c.client.Call("Server.Set", args, reply)
	if err != nil {
		log.Fatal("Set: ", err)
	}

	status := reply.Status
	return status
}

func (c *Client) Get(key string) (status StatusType, val string) {
	args := &Args{Key: key}
	reply := new(Reply)
	err := c.client.Call("Server.Get", args, reply)
	if err != nil {
		log.Fatal("Get: ", err)
	}
	return reply.Status, reply.Val
}
