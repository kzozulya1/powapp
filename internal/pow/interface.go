package pow

// POW interface of challenge-response facility
type POW interface {
	// GenerateChallenge makes random challenge
	GenerateChallenge() ([]byte, error)
	// SolveChallenge do hard-work, called by client
	SolveChallenge(challenge []byte) string
	// VerifyResponse checks client's work result - nonce
	VerifyResponse(challenge []byte, responseNonce []byte, len int) bool
}
