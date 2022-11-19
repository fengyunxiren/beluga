package global

const VERSION = `0.0.1`

const X_REQUEST_HEADER = "X-Request-Id"

const (
	CONTEXT_REQUEST_ID = "RequestId"
	CONTEXT_LOGGER     = "Logger"
	CONTEXT_DB         = "DB"
)

const (
	DIGIT     = "0123456789"
	CAPITAL   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
	SPECIAL   = `~!@#$%^&*()_+<>?,./[]{}\|`
	LETTER    = CAPITAL + LOWERCASE
)

const (
	AUTH_HEADER = "X-Auth-Token"
)
