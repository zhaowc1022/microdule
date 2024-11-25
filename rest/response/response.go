package response

const (
	ERROR     = 500
	SUCCESS   = 200
	JWTERROR  = 401
	AUTHERROR = 403
)

type (
	Response interface {
		Ok() error
		OkWithMessage(string) error
		OkWithData(interface{}) error
		OkWithString(string) error
		OkWithDetailed(interface{}, string) error
		Fail() error
		FailWithMessage(string) error
		FailWithDataMessages(interface{}) error
		FailWithDetailed(interface{}, string) error
		JWTFailWithDetailed(interface{}, string) error
		OkWithDataLogin(interface{}, string) error
		OkStatusData(int, interface{}, string) error
	}

	ResponseData struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"message"`
	}
)
