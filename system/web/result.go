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

func GenSuccess(msg string) *ResponseBean {
	return &ResponseBean{200, msg, ""}
}

func GenSuccessMsg(data interface{}) *ResponseBean {
	return &ResponseBean{200, "OK", data}
}

func GenOperationErrorMsg() *ResponseBean {
	return &ResponseBean{400, "Bad Request", ""}
}

func GenNotFondMsg() *ResponseBean {
	return &ResponseBean{404, "API Not Found", ""}
}

func GenFailedMsg(errMsg string) *ResponseBean {
	return &ResponseBean{500, errMsg, ""}
}
