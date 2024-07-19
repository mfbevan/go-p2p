package node

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	peerstore "github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	multiaddr "github.com/multiformats/go-multiaddr"
)

type Config struct {
	Port int
}

type Node struct {
	Port        int
	Node        host.Host
	PingService *ping.PingService
}

// Create a new node with the given configuration
//   - Port: The port to listen on
func New(config Config) *Node {
	return &Node{Port: config.Port}
}

// Start the node
func (n *Node) Start() {
	fmt.Println("Starting node...")
	address := fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", n.Port)
	node, err := libp2p.New(
		libp2p.ListenAddrStrings(address),
		libp2p.Ping(false), // Disables internally p2p ping service
	)

	if err != nil {
		panic(err)
	}

	n.Node = node
	n.SetupPingService()
}

// Setup the signal handler to stop the node when a SIGINT or SIGTERM is received
func (n *Node) AwaitTerminate() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
}

// Stop the node
func (n *Node) Stop() {
	n.Node.Close()
}

// Setup the ping service
func (n *Node) SetupPingService() *ping.PingService {
	pingService := &ping.PingService{Host: n.Node}
	n.Node.SetStreamHandler(ping.ID, pingService.PingHandler)

	n.PingService = pingService
	return pingService
}

// Print the information for this peer node in multiaddr format
func (n *Node) PrintInfo() {
	peerInfo := peerstore.AddrInfo{
		ID:    n.Node.ID(),
		Addrs: n.Node.Addrs(),
	}
	addrs, err := peerstore.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Node Address:", addrs[0])
}

func (n *Node) SetupPeers() {
	if len(os.Args) > 1 {
		fmt.Println("Connecting to peer", os.Args[1])
		// Get the multiaddress from the command line arguments
		addr, err := multiaddr.NewMultiaddr(os.Args[1])
		if err != nil {
			panic(err)
		}
		// Create a new peer from the multiaddress
		peer, err := peerstore.AddrInfoFromP2pAddr(addr)
		if err != nil {
			panic(err)
		}
		// Connect to the peer
		if err := n.Node.Connect(context.Background(), *peer); err != nil {
			panic(err)
		}
		// Send 5 ping messages to the peer
		fmt.Println("Sending 5 Ping Messages to ", addr)
		ch := n.PingService.Ping(context.Background(), peer.ID)
		for i := 0; i < 5; i++ {
			res := <-ch
			fmt.Println("pinged", addr, "in", res.RTT)
		}
	} else {
		n.AwaitTerminate()
	}
}
