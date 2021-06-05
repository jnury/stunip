package main

import (
	"flag"
	"fmt"
	"github.com/pion/stun"
)

func main() {
	server := flag.String("server", "stun.l.google.com:19302", "STUN server and port (server:port)")
	flag.Parse()

	// Creating a "connection" to STUN server.
	c, err := stun.Dial("udp", *server)
	if err != nil {
		panic(err)
	}
	// Building binding request with random transaction id.
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	// Sending request to STUN server, waiting for response message.
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			panic(res.Error)
		}
		// Decoding XOR-MAPPED-ADDRESS attribute from message.
		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			panic(err)
		}
		fmt.Println(xorAddr.IP)
	}); err != nil {
		panic(err)
	}
}