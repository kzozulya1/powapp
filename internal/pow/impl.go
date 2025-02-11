package pow

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	// defaultPrefix is a difficulty level (number of leading zeros)
	defaultPrefix = "000000"
	// defaultChallengeBytes is number of bytes of the challenge buffer
	defaultChallengeBytes = 32
)

// Impl is an implementation of POW interface
type Impl struct {
	prefix         string
	challengeBytes int
}

// OptFunc opt func
type OptFunc func(powImpl *Impl)

// WithPrefix overrides default prefix
func WithPrefix(prefix string) OptFunc {
	return func(p *Impl) {
		p.prefix = prefix
	}
}

// WithChallengeBytes overrides default challengeBytes
func WithChallengeBytes(challengeBytes int) OptFunc {
	return func(p *Impl) {
		p.challengeBytes = challengeBytes
	}
}

func New(opts ...OptFunc) *Impl {
	p := &Impl{
		prefix:         defaultPrefix,
		challengeBytes: defaultChallengeBytes,
	}

	// apply options if set
	for _, fn := range opts {
		fn(p)
	}

	return p
}

// GenerateChallenge creates a random byte slice to serve as a challenge.
func (w *Impl) GenerateChallenge() ([]byte, error) {
	challenge := make([]byte, w.challengeBytes)

	// fill with random data
	_, err := rand.Read(challenge)
	if err != nil {
		return nil, err
	}

	return challenge, nil
}

// VerifyResponse checks client's work result - nonce
func (w *Impl) VerifyResponse(challenge, responseNonce []byte, respLen int) bool {
	clientNonce := string(responseNonce[:respLen])

	// server verifies it before send wisdom quote
	expectedHash := sha256.Sum256(append(challenge, []byte(clientNonce)...))
	expectedHashStr := hex.EncodeToString(expectedHash[:])

	return strings.HasPrefix(expectedHashStr, w.prefix)
}

// SolveChallenge do hard-work, called by client
func (w *Impl) SolveChallenge(challenge []byte) string {
	var nonce int64
	// client must compute a hash that meets certain criteria (e.g., starts with a certain number of zeros).
	for {
		nonce++
		// skip lint error appendAssign: append result not assigned to the same slice
		//nolint:gocritic
		data := append(challenge, []byte(fmt.Sprintf("%d", nonce))...)
		hash := sha256.Sum256(data)
		hashStr := hex.EncodeToString(hash[:])
		if hashStr[:len(w.prefix)] == w.prefix {
			return fmt.Sprintf("%d", nonce)
		}
	}
}
