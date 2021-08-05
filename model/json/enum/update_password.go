package enum

type UpdatePasswordCode int

const (
	NoError UpdatePasswordCode = iota
	NotExistsUser
	ParseParamError
	NotEqualOld
	EqualOldNew
	UpdateError
)
