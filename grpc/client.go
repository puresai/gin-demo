package main

import (
    "context"
    "fmt"

    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("127.0.0.1:57501", grpc.WithInsecure())
    if err != nil {
        fmt.Println(err)
    }
    client := NewHelloServiceClient(conn)
    r, err := client.SayHello(context.Background(), &Request{Data: []byte("send from client")})
    fmt.Printf("%s\n", string(r.Data))
}