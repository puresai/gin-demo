package main

import (
    "context"
    "fmt"
    "net"
    "net/http"
	
    "google.golang.org/grpc"
)

type HelloServiceServerImpl struct {
}

func (s *HelloServiceServerImpl) SayHello(c context.Context, req *http.Request) (*http.Response, error) {
    fmt.Printf("%s\n", string(req.Data))

    resp := http.Response{}
    resp.Data = []byte("hello from server")

    return &resp, nil
}

func main() {
    lis, err := net.Listen("tcp", "127.0.0.1:57501")
    if err != nil {
        fmt.Println(err)
        return
    }
    s := grpc.NewServer()
    RegisterHelloServiceServer(s, &HelloServiceServerImpl{})
    fmt.Printf("Server listening on 127.0.0.1:57501\n")
    s.Serve(lis)
}