package main

import "github.com/mfbevan/go-p2p/node"

func main() {
	config := node.Config{}
	n := node.New(config)
	n.Start()
	n.PrintInfo()
	n.SetupPeers()
}
