package util

import (
	"net/http"
	"tdpos/common"
)

func Success(data *interface{}) *common.Response {
	return &common.Response{Code: http.StatusOK, Msg: "ok", Data: data}
}

func NotFound() *common.Response {
	return &common.Response{Code: http.StatusNotFound, Msg: "Not Found"}
}

func BadRequest() *common.Response {
	return &common.Response{Code: http.StatusBadRequest, Msg: "请求参数错误"}
}

func Error(code int, msg string) *common.Response {
	return &common.Response{Code: code, Msg: msg}
}
