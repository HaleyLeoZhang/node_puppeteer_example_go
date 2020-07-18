package constant

const (
	HTTP_CONTEXT_GET_CODE = "has_http_code"
	// 通用业务响应码
	HTTP_RESPONSE_CODE_SUCCESS          = 200
	HTTP_RESPONSE_CODE_PARAM_INVALID    = 400
	HTTP_RESPONSE_CODE_SOURCE_NOT_FOUND = 404
	HTTP_RESPONSE_CODE_GENERAL_FAIL     = 500
	HTTP_RESPONSE_CODE_TIMEOUT          = 504
	// 业务响应码
)

// 业务码对应文案
var HTTP_RESPONSE_MAP = map[int]string{
	HTTP_RESPONSE_CODE_SUCCESS:          "success",
	HTTP_RESPONSE_CODE_PARAM_INVALID:    "Param is invalid",
	HTTP_RESPONSE_CODE_SOURCE_NOT_FOUND: "Source not found",
	HTTP_RESPONSE_CODE_GENERAL_FAIL:     "Service busy",
	HTTP_RESPONSE_CODE_TIMEOUT:          "Timeout",
}
