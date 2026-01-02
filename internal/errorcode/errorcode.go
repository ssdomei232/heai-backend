package errorcode

type ErrorCode int

const (
	CodeSuccess             ErrorCode = 200   // 成功
	CodeUnAuthorized        ErrorCode = 10001 // 未登陆
	CodeInsufficientBalance ErrorCode = 10002 // 余额不足
	CodeInvalidRequest      ErrorCode = 10003 // 请求参数错误
	CodeInternalError       ErrorCode = 10004 // 系统内部错误
	CodeGenerateImageFailed ErrorCode = 10005 // 生成图片失败
	CodeUpstreamAPIError    ErrorCode = 10006 // 上游接口异常
	CodeDatabaseError       ErrorCode = 10007 // 数据库错误
)
