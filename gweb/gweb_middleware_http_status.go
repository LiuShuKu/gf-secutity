// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.

// Copyright (c)  2022 LiuShuKu
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/LiuShuKu/gf-secutity.
// Project Name :   gf-security
// Author       :   liushuku@yeah.net
// HandlerStatus*() Functions only contain status codes common in web development

// Package gweb provide common tool functions for web development.
package gweb

import (
	"gf-secutity/gresult"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

// HandlerStatus400  missing request parameter
func HandlerStatus400(request *ghttp.Request) {
	request.Middleware.Next()
	if request.Response.Status == http.StatusBadRequest {
		gresult.BadRequest(request)
	}
}

// HandlerStatus401 Request requires user authentication
func HandlerStatus401(request *ghttp.Request) {
	request.Middleware.Next()
	if request.Response.Status == http.StatusUnauthorized {
		gresult.Unauthorized(request)
	}
}

// HandlerStatus404  requested content does not exist
func HandlerStatus404(request *ghttp.Request) {
	request.Middleware.Next()
	if request.Response.Status == http.StatusNotFound {
		gresult.NotFound(request)
	}
}

// HandlerStatus405  Disable the method specified in the request。
func HandlerStatus405(request *ghttp.Request) {
	request.Middleware.Next()
	if request.Response.Status == http.StatusMethodNotAllowed {
		gresult.MethodNotAllowed(request)
	}
}

// HandlerStatus415  The requested format is not supported by the requested page。
func HandlerStatus415(request *ghttp.Request) {
	request.Middleware.Next()
	if request.Response.Status == http.StatusUnsupportedMediaType {
		gresult.UnsupportedMediaType(request)
	}
}

// HandlerStatus500
// This middleware only responds to internal service exceptions to the client ;
// It does not have data leakage problems, it is recommended to use
func HandlerStatus500(request *ghttp.Request) {
	request.Middleware.Next()
	if request.Response.Status >= http.StatusInternalServerError {
		gresult.SystemError("", request)
	}
}

// HandlerServer500
// This middleware will display the error message of the current service in the msg field
// Data leakage may occur. If the wrong data contains sensitive information such as database connection failure information,
// it will be exposed. It is not recommended for use in production environments.
func HandlerServer500(request *ghttp.Request) {
	request.Middleware.Next()
	if request.Response.Status >= http.StatusInternalServerError {
		err := request.GetError()
		if err != nil {
			gresult.SystemError(err.Error(), request)
		}
		gresult.SystemError("", request)
	}
}
