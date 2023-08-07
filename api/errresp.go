package api

import (
	"community_demo/serializer"
	"encoding/json"
	"fmt"
)

// 返回错误信息ErrorResponse
func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "Json类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}
	return serializer.Response{
		Status: 40002,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
