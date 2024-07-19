package main

import "github.com/mfbevan/go-p2p/node"

func main() {
	config := node.Config{Port: 61877}
	n := node.New(config)
	n.Start()
}
