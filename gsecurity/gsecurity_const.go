// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.

// Copyright (c)  2022 LiuShuKu
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/LiuShuKu/gf-secutity.
// Project Name :   gf-security
// Author       :   liushuku@yeah.net
// responsible for project configuration

// Package gsecurity project core modules' security verification .
package gsecurity

//  specific configuration
const (
	// profile name
	// the project will read the configuration file
	configFileName = "gf-security.yml"

	// configuration header name
	configName = "gf-security"
)

// all configurable items
const (

	// token generated token name , also the cookie name
	tokenName = "gf-token"

	// token expiration time
	timeoutName = "timeout"

	// temporary validity period of token
	activityTimeoutName = "activityTimeout"

	// Whether to allow the same account to log in concurrently
	isConcurrentName = "isConcurrent"

	// whether to share a token
	isShareName = "isShare"

	// maximum number of logins for the same account
	maxLoginCountName = "maxLoginCount"

	// Whether to try to read the Token from the request body
	isReadBodyName = "isReadBody"

	// whether to try to read the token from the header
	isReadHeadName = "isReadHead"

	// whether to try to read the token from the cookie
	isReadCookieName = "isReadCookie"

	// token style
	tokenStyleName = "tokenStyle"

	// By default, the data is stored in the local cache, and the interval between expired data is cleaned up each time.
	dataRefreshPeriodName = "dataRefreshPeriod"

	// whether to log in when getting token session
	tokenSessionCheckLoginName = "tokenSessionCheckLogin"

	// whether to turn on automatic renewal
	autoRenewName = "autoRenew"

	// token prefix
	tokenPrefixName = "tokenPrefix"

	// Whether to print version glyphs when initializing the configuration
	isPrintName = "isPrint"

	// whether to print the operation log
	isLogName = "isLog"

	// jwtKey
	jwtSecretKeyName = "jwtSecretKey"

	// validity period of id token
	idTokenTimeoutName = "idTokenTimeout"

	// Http Basic authenticated account and password
	basicName = "basic"

	// Configure the network access address of the current project
	currDomainName = "currDomain"

	// check whether Id-Token
	checkIdTokenName = "checkIdToken"
)
