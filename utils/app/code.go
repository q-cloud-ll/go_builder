package app

type ResCode int64

const (
	// CodeSuccess 成功（默认返回状态码）
	CodeSuccess ResCode = 0
	// CodeSeverError 全局未知异常
	CodeSeverError ResCode = 500
	// CodeBadRequest 请求失败（一般前端处理，不常用）
	CodeBadRequest ResCode = 400
	// CodeDataNotFount 请求资源不存在（静态资源不存在，不常用）
	CodeDataNotFount ResCode = 404
	// CodeLoginExpire 登录认证异常
	CodeLoginExpire ResCode = 401
)

/*
通用业务
*/
const (
	/*
	   1001-1010 通用操作相关
	*/
	// CodeOperationFail 操作失败
	CodeOperationFail ResCode = 1001 + iota
	// CodeSelectOperationFail 查询操作失败
	CodeSelectOperationFail
	// CodeUpdateOperationFail 更新操作失败
	CodeUpdateOperationFail
	// CodeDeleteOperationFail 删除操作失败
	CodeDeleteOperationFail
	// CodeInsertOperationFail 新增操作失败
	CodeInsertOperationFail
	CodeInvalidParam

	/*
	   1011-1050 例如登录注册相关
	*/

)

/*
   -----------go_api 业务相关（2xxx）------------
*/

var codeMsgMap = map[ResCode]string{
	CodeSuccess:             "success",
	CodeSeverError:          "服务器繁忙请重试",
	CodeBadRequest:          "请求失败",
	CodeDataNotFount:        "未找到资源",
	CodeLoginExpire:         "请登录后重试",
	CodeOperationFail:       "操作失败",
	CodeSelectOperationFail: "查询操作失败！",
	CodeUpdateOperationFail: "更新操作失败！",
	CodeDeleteOperationFail: "删除操作失败！",
	CodeInsertOperationFail: "新增操作失败！",
	CodeInvalidParam:        "请求参数错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeSeverError]
	}
	return msg
}
