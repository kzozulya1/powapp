package verification

import (
	"github.com/golang/mock/gomock"
	"github.com/kzozulya1/powapp/internal/pow"
)

// mocker is a model for tests
type mocker struct {
	pow  *pow.MockPOW
	conn *MockConn
}

// newMocker creates new mocker
func newMocker(ctrl *gomock.Controller) *mocker {
	return &mocker{
		pow:  pow.NewMockPOW(ctrl),
		conn: NewMockConn(ctrl),
	}
}
