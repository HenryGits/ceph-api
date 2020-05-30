package web

/**
 封装请求返回值
结构体中的声明变量首字母必须大写 不然无法被beego解析
*/
type ResponseBean struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS    = 200
	NotFount   = 404
	ERROR      = 500
	BadRequest = 400
)

func GenSuccess(msg string) *ResponseBean {
	return &ResponseBean{SUCCESS, msg, ""}
}

func GenSuccessMsg(data interface{}) *ResponseBean {
	return &ResponseBean{SUCCESS, "OK", data}
}

func GenOperationErrorMsg() *ResponseBean {
	return &ResponseBean{BadRequest, "Bad Request", ""}
}

func GenNotFondMsg() *ResponseBean {
	return &ResponseBean{NotFount, "API Not Found", ""}
}

func GenFailedMsg(errMsg string) *ResponseBean {
	return &ResponseBean{ERROR, errMsg, ""}
}
