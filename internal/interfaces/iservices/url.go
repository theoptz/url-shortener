package iservices

var _ URLValidatorService = &MockURLValidatorService{}

//go:generate mockery --all --inpackage --case underscore
type URLValidatorService interface {
	Validate(string) bool
}
