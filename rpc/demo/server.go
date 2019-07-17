package main

import (
	"net/rpc"
	"net/http"
	"log"
)

type Params struct {
	Width, Height int;
}

type Rect struct {}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height;
	return nil;
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2;
	return nil;
}

func main() {
	rect := new(Rect);
	rpc.Register(rect);
	rpc.HandleHTTP();
	err := http.ListenAndServe(":8099", nil);
	if err != nil {
		log.Fatal(err)
	}
}