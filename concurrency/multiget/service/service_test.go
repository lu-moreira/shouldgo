package service_test

import (
	"context"
	"testing"

	"github.com/lu-moreira/shouldgo/concurrency/multiget/dto"
	"github.com/lu-moreira/shouldgo/concurrency/multiget/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func Test_WhenGETOperationsReturns200OK_ShouldReturnAError(t *testing.T) {
	defer goleak.VerifyNone(t)

	response, err := service.GetUserInformation(context.TODO(), dto.GetUserRequest{
		GrootID: 21947,
	})
	assert.Nil(t, err)
	require.NotNil(t, response)
	assert.NotNil(t, response.A)
	assert.NotNil(t, response.P)
	assert.NotNil(t, response.U)
}
