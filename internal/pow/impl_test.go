package pow

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPOWImpl(t *testing.T) {
	tests := []struct {
		name  string
		setup func(pow POW, challenge, nonce []byte, nonceLen *int)
		want  bool
	}{
		{
			name: "challenge-solved",
			setup: func(pow POW, challenge, nonce []byte, nonceLen *int) {
				solvedNonce := pow.SolveChallenge(challenge)
				copy(nonce, solvedNonce)
				*nonceLen = len(solvedNonce)
			},
			want: true,
		},
		{
			name: "challenge-failed",
			setup: func(_ POW, challenge, nonce []byte, nonceLen *int) {
				// emulate wrong nonce from client
				copy(nonce, []byte{52, 52, 52, 52, 52}) // 44444 as string
				*nonceLen = 5
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			powImpl := New(WithPrefix("0000")) // less zeros -> faster

			challenge, err := powImpl.GenerateChallenge()
			if err != nil {
				t.Errorf("powImpl GenerateChallenge() error = %v", err)
			}

			// create buf
			buf := make([]byte, 16)
			nonceLen := 0
			// fill buf with nonce
			tt.setup(powImpl, challenge, buf, &nonceLen)
			// check computed nonce
			verifyResult := powImpl.VerifyResponse(challenge, buf, nonceLen)

			if !gomock.Eq(verifyResult).Matches(tt.want) {
				t.Errorf("unexpected value: actual: %t, expected: %t",
					verifyResult,
					tt.want,
				)
			}
		})
	}
}
