package main

import (
	"fmt"
	"net"
	"powapp/internal/pow"
	"powapp/internal/verification"
	"powapp/internal/wisdomquotes"
)

// handleConnection handles new TCP connection
func handleConnection(conn net.Conn, powImpl pow.POW) {
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("handle connection: close conn:", err)
		}
	}()

	var err error

	// initiate challenge=response protocol and get nonce verify result
	verified, err := verification.Verify(conn, powImpl)
	if err != nil {
		fmt.Println("handle connection: pow verification:", err)
		return
	}

	result := "Invalid proof of work\n"
	if verified {
		result = wisdomquotes.Quote() + "\n"
	}

	if _, err = conn.Write([]byte(result)); err != nil {
		fmt.Println("handle connection: write challenge result:", err)
	}
}
