package iservices

var _ CounterService = &MockCounterService{}

//go:generate mockery --all --inpackage --case underscore
type CounterService interface {
	Inc() (int64, error)
}
