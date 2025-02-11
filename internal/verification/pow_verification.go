package verification

import (
	"errors"
	"net"
	"powapp/internal/pow"
)

// maxNonceSize max buffer size for client nonce response
const maxNonceSize = 16

// Verify generates a challenge, sends it to the client,
// waits for client's response (the nonce).
// Checks if the nonce solves the challenge.
func Verify(conn net.Conn, powImpl pow.POW) (bool, error) {
	// generate new challenge
	challenge, err := powImpl.GenerateChallenge()
	if err != nil {
		return false, errors.New("generating challenge:" + err.Error())
	}

	// send the challenge to the client
	_, err = conn.Write(challenge)
	if err != nil {
		return false, errors.New("conn write challenge:" + err.Error())
	}

	// read the client's response (nonce)
	buf := make([]byte, maxNonceSize)

	// server waits for client's computation result limited time (read timeout)
	n, err := conn.Read(buf)
	if err != nil {
		return false, errors.New("conn nonce read:" + err.Error())
	}

	// then verify nonce
	if powImpl.VerifyResponse(challenge, buf, n) {
		return true, nil
	}

	return false, nil
}
