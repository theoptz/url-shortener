package iservices

var _ CodeService = &MockCodeService{}

//go:generate mockery --all --inpackage --case underscore
type CodeService interface {
	Encode(int64) string
	Decode(string) int64
}
