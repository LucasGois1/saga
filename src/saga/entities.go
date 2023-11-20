package saga

import (
	"errors"
	"log"
)

//go:generate go run github.com/golang/mock/mockgen -package saga_test -destination=./mock_saga_test.go -source=./entities.go

var ErrSagaAlreadyFinished = errors.New("can't walk forward because this saga is already finished")


// Possible saga states
type Status string

const (
	STARTED Status = "STARTED"
	COMPENSATED Status = "COMPENSATED"
	FINISHED Status = "FINISHED"
)

type Saga interface {
	Identifier() string
	Execute() error
	Compensate() error
	OnStep(step *Step)
	next() *Step
	walkForward() error
	walkBackward() error
	isFinished() bool
	CurrentStatus() Status
}

type Step struct {
	Name string
	Next *Step
	Do   func() error
	Compensation func() error
}

type SimpleSaga struct {
	identifier string
	steps 	[]*Step
	actualStepIndex uint8
	status Status
}

func newSimpleSaga(identifier string) *SimpleSaga {
	return &SimpleSaga{
		identifier: identifier,
		steps: []*Step{},
		actualStepIndex: 0,
		status: STARTED,
	}
}

func(s *SimpleSaga) next() *Step {
	return s.steps[s.actualStepIndex + 1]
}

func (s *SimpleSaga) Identifier() string {
	return s.identifier
}

func (s *SimpleSaga) Execute() error {
	err := s.next().Do()

	if err != nil {
		return err
	}

	s.walkForward()

	return nil
}


func (s *SimpleSaga) Compensate() error {
	for s.next() != nil {
		err := s.next().Compensation()

		if err != nil {
			return err
		}

		s.walkBackward()
	}

	s.status = COMPENSATED

	return nil
}


func (s *SimpleSaga) walkForward() error {
	if s.actualStepIndex == uint8(len(s.steps) - 1) {
		return ErrSagaAlreadyFinished
	}

	s.actualStepIndex++

	return nil
}

func (s *SimpleSaga) walkBackward() error {
	if s.actualStepIndex == 0 {
		return errors.New("can't walk backwards because it's the first step")
	}

	s.actualStepIndex--

	return nil
}

func (s *SimpleSaga) isFinished() bool {
	err := s.walkForward()

	if err != nil { 
		if err == ErrSagaAlreadyFinished {
			s.status = FINISHED
			return true
		} else {
			log.Fatalf("Unexpected error while checking if saga have been finished: %v", err)
		}
	}

	return false
}

func (s *SimpleSaga) OnStep(step *Step) {
	s.steps = append(s.steps, step)
}

func (s *SimpleSaga) CurrentStatus() Status {
	return s.status
}

type SimpleSagaBuilder struct {
	saga Saga
}

func NewSimpleSagaBuilder(identifier string) *SimpleSagaBuilder {
	return &SimpleSagaBuilder{
		saga: newSimpleSaga(identifier),
	}
}

func (s *SimpleSagaBuilder) Build() Saga {
	return s.saga
}


type IBuilder interface {
	Build() Saga
}

type SagaBuilder interface {
	IBuilder
}

type SagaDirector struct {
	builder SagaBuilder
}

func NewSagaDirector(b SagaBuilder) *SagaDirector {
	return &SagaDirector{
		builder: b,
	}
}

func (s *SagaDirector) BuildSaga() Saga {
	return s.builder.Build()
}
