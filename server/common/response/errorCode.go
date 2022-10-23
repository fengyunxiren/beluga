package response

type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewErrorCode(code int, msg string) ErrorCode {
	return ErrorCode{
		Code: code,
		Msg:  msg,
	}
}

var (
	OK                      = NewErrorCode(0, "success")
	ERR_UNKNOWN             = NewErrorCode(1, "unknown")
	ERR_BAD_REQUEST         = NewErrorCode(400001, "Request parameter is not match")
	ERR_NOT_ALLOWED         = NewErrorCode(405001, "The method is not allowed")
	ERR_SERVER_500          = NewErrorCode(500001, "The server encountered an unknown error")
	ERR_SERVER_DB_NOT_FOUND = NewErrorCode(500101, "The server encountered an unknown error")
	// View error
	ERR_SERVER_VIEW_CREATE = NewErrorCode(600001, "Create model failed")
	ERR_SERVER_VIEW_GET    = NewErrorCode(600002, "Get model failed")
	ERR_SERVER_VIEW_UPDATE = NewErrorCode(600003, "Update model failed")
	ERR_SERVER_VIEW_DELETE = NewErrorCode(600004, "Delete model failed")
)
