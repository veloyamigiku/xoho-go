package enum

type UpdatePasswordCode int

const (
	NotEqualOld UpdatePasswordCode = iota
	NotEqualOldNew
)
