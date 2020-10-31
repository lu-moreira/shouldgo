package channels_test

import (
	"testing"

	"github.com/lu-moreira/shouldgo/channels"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func Test_StartNto1(t *testing.T) {
	defer goleak.VerifyNone(t)
	assert.NotPanics(t, func() {
		channels.StartNto1()
	})
}
