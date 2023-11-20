package saga

//go:generate mockgen -destination mock_saga.go -package saga_test github.com/LucasGois1/saga Saga


type Saga interface {
	Execute() error
	Compensate() error
}

type Step struct {
	Name string
	Do   func() error
	Undo func() error
}

type SimpleSaga struct {
	Steps []Step
}

type IBuilder interface {
	build() *Saga
}

type SagaBuilder interface {
	IBuilder
}

type SagaDirector struct {
	builder *SagaBuilder
}

func NewSagaDirector(b *SagaBuilder) *SagaDirector {
	return &SagaDirector{
		builder: b,
	}
}
