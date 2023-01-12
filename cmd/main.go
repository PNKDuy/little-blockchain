package main

import (
	"context"
	"fmt"

	p2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
)

func main() {
	host1, err := p2p.New()
	if err != nil {
		fmt.Println(err)
	}

	defer host1.Close()

	fmt.Printf("Host1 ID is %s\n", host1.ID())

	// private key
	priv, _, err := crypto.GenerateKeyPair(
		crypto.Ed25519,
		-1,
	)

	if err != nil {
		fmt.Println(err)
	}

	host2, err := p2p.New(
		p2p.Identity(priv),
	)

	fmt.Printf("Host2 ID is %s\n", host2.ID())

	err = host1.Connect(context.Background(), *host.InfoFromHost(host2))
	if err != nil {
		panic(err)
	}

	fmt.Println(host2.ID(), "is connected via ", host1.ID())
}
