package request

type Request interface {
	GetVal(c any) error
	GetVerifyVal(c any) error
}
