// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
// Project Name :   gf-security
// Author       :   liushuku@yeah.net
package gresult

import "github.com/gogf/gf/v2/net/ghttp"

// Success respond to success messages only
// http status 200
func Success(request *ghttp.Request) {
	Response(successCode, nil, successMsg, request)
}

// SucceedData Data in response to successful operation
func SuccessData(data interface{}, request *ghttp.Request) {
	Response(successCode, data, successMsg, request)
}

// SuccessDataTotal Response success data and total
// The total parameter of this method will be equal to data
func SuccessDataTotal(data interface{}, total int64, request *ghttp.Request) {
	ResponseTotal(successCode, data, successMsg, total, false, request)
}

// SuccessDataTotal Response success data and total
// This method total will be inside data
func SuccessDataTotalInner(data interface{}, total int64, request *ghttp.Request) {
	ResponseTotal(successCode, data, successMsg, total, true, request)
}
