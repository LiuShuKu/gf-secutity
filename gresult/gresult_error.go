// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.

// Copyright (c)  2022 LiuShuKu
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/LiuShuKu/gf-secutity.
// Project Name :   gf-security
// Author       :   liushuku@yeah.net

package gresult

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// Fail respond to error messages
// for business errors
func Fail(msg string, request *ghttp.Request) {
	Response(failCode, nil, msg, request)
}

// ServerError response service error
// TODO 保留
func ServerError(err error, request *ghttp.Request) {
}

// BadRequest missing request parameter
// status:400
func BadRequest(request *ghttp.Request) {
	SystemError(paramFailMsg, request)
}

// Unauthorized Request requires user authentication
// status:401
func Unauthorized(request *ghttp.Request) {
	Response(authFailCode, nil, authFailMsg, request)
}

// NotFound  requested content does not exist
// status:404
func NotFound(request *ghttp.Request) {
	SystemError(notFoundFailMsg, request)
}

// MethodNotAllowed Disable the method specified in the request。
// status:405
func MethodNotAllowed(request *ghttp.Request) {
	SystemError(methodFailMsg, request)
}

// UnsupportedMediaType The requested format is not supported by the requested page。
// status:415
func UnsupportedMediaType(request *ghttp.Request) {
	SystemError(mediaFailMsg, request)
}

// SystemError Internal server error
// status:500
func SystemError(msg string, request *ghttp.Request) {
	if msg == "" {
		msg = errorMsg
	}
	Response(errorCode, nil, msg, request)
}
