package main

import (
    "fmt"
)

type user struct {
    name string
}

type Admin struct {
    user
}

func main() {
    var ad Admin
    ad.name="张三"
    fmt.Println(ad)
}