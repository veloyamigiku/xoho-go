package enum

type UpdatePasswordCode int

const (
	NoError UpdatePasswordCode = iota
	ParseParamError
	NotEqualOld
	EqualOldNew
)
