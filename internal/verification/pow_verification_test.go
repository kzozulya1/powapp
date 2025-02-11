package verification

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

var challenge = []byte("challenge1")

// TestPOWVerification tests Verify route - for handling new TCP conns
func TestPOWVerification(t *testing.T) {
	tests := [...]struct {
		name    string
		setup   func(ctrl *gomock.Controller, mm *mocker)
		want    bool
		wantErr bool
	}{
		{
			name: "positive-challenge-solved",
			setup: func(ctrl *gomock.Controller, mm *mocker) {
				// we can write test next way: mock args from testing routine to exactly match of
				// expected types

				// nonce - emulate response from client
				nonce := []byte{52, 57, 53, 50, 57} // 46529 as string
				buf := make([]byte, maxNonceSize)

				mm.pow.EXPECT().GenerateChallenge().Return(challenge, nil)
				mm.conn.EXPECT().Write(challenge).Return(len(challenge), nil)
				// setup read data in caller with nonce
				mm.conn.EXPECT().Read(buf).
					Do(func(b []byte) {
						copy(b, nonce)
					}).
					Return(len(challenge), nil)

				// prepare buffer for verification
				bufForVerification := make([]byte, maxNonceSize)
				copy(bufForVerification, nonce)
				mm.pow.EXPECT().VerifyResponse(challenge, bufForVerification, len(challenge)).Return(true)
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "negative-challenge-solved",
			setup: func(ctrl *gomock.Controller, mm *mocker) {
				// but also we can write quick test: mock args in most abstract way, mock only result
				// of crucial func VerifyResponse
				mm.pow.EXPECT().GenerateChallenge().Return(challenge, nil)
				mm.conn.EXPECT().Write(gomock.Any()).Return(len(challenge), nil)
				mm.conn.EXPECT().Read(gomock.Any()).Return(len(challenge), nil)
				mm.pow.EXPECT().VerifyResponse(challenge, gomock.Any(), len(challenge)).Return(false)
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "negative-conn-write-failed",
			setup: func(ctrl *gomock.Controller, mm *mocker) {
				mm.pow.EXPECT().GenerateChallenge().Return(challenge, nil)
				mm.conn.EXPECT().Write(gomock.Any()).Return(0, errors.New("err while conn write"))
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "negative-conn-read-failed",
			setup: func(ctrl *gomock.Controller, mm *mocker) {
				mm.pow.EXPECT().GenerateChallenge().Return(challenge, nil)
				mm.conn.EXPECT().Write(gomock.Any()).Return(len(challenge), nil)
				mm.conn.EXPECT().Read(gomock.Any()).Return(0, errors.New("err while conn read"))
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mocker := newMocker(ctrl)
			tt.setup(ctrl, mocker)

			verifyResult, err := Verify(mocker.conn, mocker.pow)

			if (err != nil) != tt.wantErr {
				t.Errorf("pow verification: error = %v, wantErr %v", err, tt.wantErr)
			}

			if !gomock.Eq(verifyResult).Matches(tt.want) {
				t.Errorf("unexpected value: actual: %t, expected: %t",
					verifyResult,
					tt.want,
				)
			}
		})
	}
}
