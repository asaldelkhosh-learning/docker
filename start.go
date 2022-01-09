package main

import "github.com/hamed-yousefi/gowl"

func main() {
	pool := gowl.NewPool(4)
	defer pool.Close()
	_ = pool.Start()
}
