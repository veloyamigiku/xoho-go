package enum

type ResignCode int

const (
	ResignCodeNoError ResignCode = iota
	ResignCodeNotExistsUser
	ResignCodeParseParamError
	ResignCodeDbError
)
