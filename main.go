package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    sampler := make_sampler()
    fmt.Println(sampler)
}
