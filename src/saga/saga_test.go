package saga_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"saga/src/saga"
)

func ShouldCreateASagaDirector(t *testing.T) {
	ctlr := gomock.NewController(t)

	defer ctlr.Finish()

	result := saga.NewSagaDirector()

	assert.NotNil(result)
}