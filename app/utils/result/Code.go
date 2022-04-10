package result

var (
	SUCCESS = response(200, "请求成功") // 通用成功
	FAIL    = response(100, "")     // 通用错误
)
