// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.

// Copyright (c)  2022 LiuShuKu
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/LiuShuKu/gf-secutity.
// Project Name :   gf-security
// Author       :   liushuku@yeah.net
// It is not recommended operating the functions in gresult.go,
// you should use the functions in gresult_error.go and gresult_success.go to operate

// Package gresult Provides a simple data return format;
package gresult

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

// custom response structure
type result struct {
	Code int         `json:"code"` // Response code
	Msg  string      `json:"msg"`  // Response chinese message
	Data interface{} `json:"data"` // Response data
}

// custom response structure
type resultTotal struct {
	Code  int         `json:"code"`  // Response code
	Msg   string      `json:"msg"`   // Response chinese message
	Data  interface{} `json:"data"`  // Response data
	ToTal int64       `json:"total"` // Response total
}

// code = 0 all success
// code = 1 Identify business errors, foreseeable exceptions, such as account password errors
// code = 2 Identify system exceptions, unpredictable, such as 404 500
// code = -1 Identity authentication failed, authentication not completed
const (
	// authentication failed
	authFailCode = -1

	// Successful operation
	successCode = 0

	// operation failed
	failCode = 1

	// System internal exception
	errorCode = 2
)

// custom response msg
var (

	// Success
	// http status 200
	successMsg = http.StatusText(http.StatusOK)

	// The syntax of the client request is incorrect, the server cannot understand it, and the request parameters are incorrect.
	// http status 400
	paramFailMsg = http.StatusText(http.StatusBadRequest)

	// Request requires user authentication
	// http status 401
	authFailMsg = http.StatusText(http.StatusUnauthorized)

	// The requested content does not exist
	// http status 404
	notFoundFailMsg = http.StatusText(http.StatusNotFound)

	// Disable the method specified in the request。
	// http status 405
	methodFailMsg = http.StatusText(http.StatusMethodNotAllowed)

	// The requested format is not supported by the requested page。
	// http status 415
	mediaFailMsg = http.StatusText(http.StatusUnsupportedMediaType)

	// Internal server error
	// http status 500
	errorMsg = http.StatusText(http.StatusInternalServerError)
)

// Response Encapsulate Json parameter return
func Response(code int, data interface{}, msg string, request *ghttp.Request) {
	// reset response status code
	{
		resetHttpStatus(request)
	}
	request.Response.WriteJson(result{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	request.Exit()
}

// ResponseTotal Carry pagination total return
func ResponseTotal(code int, data interface{}, msg string, total int64, isInner bool, request *ghttp.Request) {
	// reset response status code
	{
		resetHttpStatus(request)
	}

	// total will be inside data
	if isInner {
		m := make(map[string]interface{}, 2)

		r := result{
			Code: code,
			Msg:  msg,
			Data: m,
		}
		{
			m["rows"] = data
			m["total"] = total

		}
		request.Response.WriteJson(r)

	} else {
		// total will be equal to data
		r := resultTotal{
			Code:  code,
			Msg:   msg,
			Data:  data,
			ToTal: total,
		}
		request.Response.WriteJson(r)
	}
	request.Exit()
}

// resetHttpStatus unified http response status code is 200
func resetHttpStatus(request *ghttp.Request) {
	request.Response.ClearBuffer()
	request.Response.Status = http.StatusOK
}
