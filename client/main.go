package main

import (
	"flag"
	"fmt"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/kzozulya1/powapp/internal/pow"
	"log"
	"net"
	"os"
)

const (
	serverAddress = "localhost:8080"
	quoteMaxSize  = 1024
)

func main() {
	serverAddrPtr := flag.String("addr", serverAddress, "address of pow server")
	flag.Parse()

	// connect to server
	fmt.Printf("Connecting to %s\n", *serverAddrPtr)
	conn, err := net.Dial("tcp", *serverAddrPtr)
	if err != nil {
		log.Println("server net dial:", err)
		return
	}
	defer func() {
		// don't forget to close conn
		if err := conn.Close(); err != nil {
			fmt.Println("close conn:", err)
		}
	}()

	// make a buffer and fill it with challenge string from server
	challenge := make([]byte, pow.CommonChallengeBytes)
	n, err := conn.Read(challenge)
	if err != nil {
		log.Println("conn read challenge:", err)
		return
	}

	// instantiate pow facility
	powImpl := pow.New(pow.WithPrefix(pow.CommonPrefix))

	// show loader-spinner ...
	fmt.Println("Please wait while solving challenge...")
	w := wow.New(os.Stdout, spin.Get(spin.Arrow3), "  ")
	w.Start()

	// do solve challenge
	solution := powImpl.SolveChallenge(challenge[:n])

	// send found solution to server
	_, err = conn.Write([]byte(solution))
	if err != nil {
		log.Println("conn write solution:", err)
		return
	}

	// read server verification result
	response := make([]byte, quoteMaxSize)
	n, err = conn.Read(response)
	if err != nil {
		log.Println("conn read server response:", err)
		return
	}

	// show server verification result
	fmt.Println(string(response[:n]))
}
