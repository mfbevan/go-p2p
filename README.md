# 🛜 Go p2p

Go implementation of peer to peer network node using [go-libp2p](https://docs.libp2p.io/guides/getting-started/go/)

## Installation

### Install Go

* Install Go and Ensure your Go version is at least 1.20.
* You can install a recent version of Go by following the [official installation instructions](https://golang.org/doc/install).
* Once installed, you should be able to run go version and see a version >= 1.20, for example:


## Run p2p nodes

> The following will demonstrate starting two nodes that will play a game of ping pong with one another.

Each node will automatically assign its own port on your local loopback interface (127.0.0.1). To run your first node, use

```bash
go run main.go
```

This will start the node and log the address of the node

See output: 

```bash
Starting node...
Node Address: /ip4/127.0.0.1/tcp/63993/p2p/12D3KooWQLLY8YyRKbhMX5r5FWQUXc2LGiocEMS67w8s63UKKaqR
```

Copy the address and open a new terminal. Run another node with the node address as an argument

```bash
go run main.go /ip4/127.0.0.1/tcp/63993/p2p/12D3KooWQLLY8YyRKbhMX5r5FWQUXc2LGiocEMS67w8s63UKKaqR
```

This will connect both nodes and send 5 consecutive ping messages to the first node, logging their output, before terminating

See output: 

```bash
Starting node...
Node Address: /ip4/127.0.0.1/tcp/64018/p2p/12D3KooWEtqY7hUDAUjtR6hy6ktF3kfrHYH1a7KLyq6kVhUyfUaG
Connecting to peer /ip4/127.0.0.1/tcp/63993/p2p/12D3KooWQLLY8YyRKbhMX5r5FWQUXc2LGiocEMS67w8s63UKKaqR
Sending 5 Ping Messages to  /ip4/127.0.0.1/tcp/63993/p2p/12D3KooWQLLY8YyRKbhMX5r5FWQUXc2LGiocEMS67w8s63UKKaqR
pinged /ip4/127.0.0.1/tcp/63993/p2p/12D3KooWQLLY8YyRKbhMX5r5FWQUXc2LGiocEMS67w8s63UKKaqR in 101.084µs
pinged /ip4/127.0.0.1/tcp/63993/p2p/12D3KooWQLLY8YyRKbhMX5r5FWQUXc2LGiocEMS67w8s63UKKaqR in 49.875µs
pinged /ip4/127.0.0.1/tcp/63993/p2p/12D3KooWQLLY8YyRKbhMX5r5FWQUXc2LGiocEMS67w8s63UKKaqR in 55.208µs
pinged /ip4/127.0.0.1/tcp/63993/p2p/12D3KooWQLLY8YyRKbhMX5r5FWQUXc2LGiocEMS67w8s63UKKaqR in 63.875µs
pinged /ip4/127.0.0.1/tcp/63993/p2p/12D3KooWQLLY8YyRKbhMX5r5FWQUXc2LGiocEMS67w8s63UKKaqR in 55.708µs
```