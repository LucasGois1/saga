package saga_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	saga "github.com/LucasGois1/saga-pattern/src/saga"
)

func TestShouldCreateASagaDirector(t *testing.T) {
	ctlr := gomock.NewController(t)

	defer ctlr.Finish()

	mockSagaBuilder := NewMockSagaBuilder(ctlr)

	sagaDirector := saga.NewSagaDirector(mockSagaBuilder)

	assert.NotNil(t, sagaDirector)
}

func TestShouldBuildASaga(t *testing.T) {
	ctlr := gomock.NewController(t)

	defer ctlr.Finish()

	mockSaga := NewMockSaga(ctlr)
	mockSaga.EXPECT().Identifier().Return("mockSaga").Times(1)

	mockSagaBuilder := NewMockSagaBuilder(ctlr)
	mockSagaBuilder.EXPECT().Build().Return(mockSaga).Times(1)

	result := saga.NewSagaDirector(mockSagaBuilder).BuildSaga()

	assert.NotNil(t, result)
	assert.Equal(t, "mockSaga", result.Identifier())
}