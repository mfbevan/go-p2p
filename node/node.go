package node

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
)

type Config struct {
	Port int
}

type Node struct {
	Port int
	Node host.Host
}

// Create a new node with the given configuration
//   - Port: The port to listen on
func New(config Config) *Node {
	return &Node{Port: config.Port}
}

// Start the node
func (n *Node) Start() {
	address := fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", n.Port)
	node, err := libp2p.New(
		libp2p.ListenAddrStrings(address),
		libp2p.Ping(false), // Disables internally p2p ping service
	)

	if err != nil {
		panic(err)
	}

	n.Node = node

	fmt.Println("Listen addresses:", node.Addrs())

	n.setupTerminator()
}

// Setup the signal handler to stop the node
func (n *Node) setupTerminator() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")

	if err := n.Node.Close(); err != nil {
		panic(err)
	}
}

// Stop the node
func (n *Node) Stop() {
	n.Node.Close()
}

// Setup the ping service
func (n *Node) SetupPingService() {
	pingService := &ping.PingService{Host: n.Node}
	n.Node.SetStreamHandler(ping.ID, pingService.PingHandler)
}
