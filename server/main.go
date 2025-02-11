package main

import (
	"fmt"
	"github.com/kzozulya1/powapp/internal/pow"
	"log"
	"net"
	"os"
	"time"
)

const (
	envTCPPort     = "POW_TCP_PORT"
	defaultTCPPort = "8080"

	envReadTimeout                 = "POW_READ_TIMEOUT"
	defaultConnReadTimeoutDuration = "5s"
)

func main() {
	// server configuration
	tcpAddr := ":" + getTCPPort()
	readTimeoutDur, err := getReadTimeoutDuration()
	if err != nil {
		fmt.Println("parse read timeout duration:", err)
		return
	}

	//nolint:gosec
	listener, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		log.Fatalln("starting server:", err)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			log.Println("closing listener:", err)
		}
	}()

	powImpl := pow.New(
		pow.WithPrefix(pow.CommonPrefix),
		pow.WithChallengeBytes(pow.CommonChallengeBytes),
	)

	fmt.Println("Server listening on ", tcpAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accepting connection:", err)
			continue
		}

		err = conn.SetReadDeadline(time.Now().Add(readTimeoutDur))
		if err != nil {
			fmt.Println("connection set read deadline:", err)
			continue
		}

		go handleConnection(conn, powImpl)
	}
}

// getTCPPort returns env tcp port or default value
func getTCPPort() string {
	if envTCPPort := os.Getenv(envTCPPort); envTCPPort != "" {
		return envTCPPort
	}

	return defaultTCPPort
}

// getReadTimeoutDuration returns env read timeout of default value
func getReadTimeoutDuration() (time.Duration, error) {
	rt := defaultConnReadTimeoutDuration
	if envReadTimeout := os.Getenv(envReadTimeout); envReadTimeout != "" {
		rt = envReadTimeout
	}

	return time.ParseDuration(rt)
}
