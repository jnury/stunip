package main

import (
	"flag"
	"fmt"
	"github.com/pion/stun"
	"os"
)

func main() {
	server := flag.String("server", "stun.l.google.com:19302", "STUN server:port")
	flag.Parse()

	// Creating a "connection" to STUN server.
	c, err := stun.Dial("udp", *server)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Building binding request with random transaction id.
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	// Sending request to STUN server, waiting for response message.
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			fmt.Println("Error:", res.Error)
			os.Exit(1)
		}
		// Decoding XOR-MAPPED-ADDRESS attribute from message.
		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(xorAddr.IP)
	}); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
