// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.

// Copyright (c)  2022 LiuShuKu
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/LiuShuKu/gf-secutity.
// Project Name :   gf-security
// Author       :   liushuku@yeah.net

// Package gsecurity .
package gsecurity

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

type Config struct{}

// GetTokenNameValue Get the data value of the field 'tokenName' in the gf-security configuration file
func (c *Config) GetTokenNameValue() (tokenValue string) {
	v := handlerConfigNotFound(join(configName, tokenName))
	tokenValue = tokenDefValue
	if v != nil {
		tokenValue = v.String()
	}
	return
}

// GetTimeoutNameValue Get the data value of the field 'timeout' in the gf-security configuration file
func (c *Config) GetTimeoutNameValue() (timeoutValue int64) {
	v := handlerConfigNotFound(join(configName, timeoutName))
	timeoutValue = timeoutDefValue
	if v != nil {
		timeoutValue = v.Int64()
	}
	return
}

// GetActivityTimeoutNameValue Get the data value of the field 'activityTimeout' in the gf-security configuration file
func (c *Config) GetActivityTimeoutNameValue() (activityTimeoutValue int64) {
	v := handlerConfigNotFound(join(configName, activityTimeoutName))
	activityTimeoutValue = activityTimeoutDefValue
	if v != nil {
		activityTimeoutValue = v.Int64()
	}
	return
}

// GetIConcurrentNameValue Get the data value of the field 'isConcurrent' in the gf-security configuration file
func (c *Config) GetIConcurrentNameValue() (isConcurrentValue bool) {
	v := handlerConfigNotFound(join(configName, isConcurrentName))
	isConcurrentValue = isConcurrentDefValue
	if v != nil {
		isConcurrentValue = v.Bool()
	}
	return
}

// GetIsShareNameValue Get the data value of the field 'isShare' in the gf-security configuration file
func (c *Config) GetIsShareNameValue() (isShareValue bool) {
	v := handlerConfigNotFound(join(configName, isShareName))
	isShareValue = isShareDefValue
	if v != nil {
		isShareValue = v.Bool()
	}
	return
}

// GetMaxLoginCountNameValue Get the data value of the field 'maxLoginCount' in the gf-security configuration file
func (c *Config) GetMaxLoginCountNameValue() (maxLoginCountValue int) {
	v := handlerConfigNotFound(join(configName, maxLoginCountName))
	maxLoginCountValue = maxLoginCountDefValue
	if v != nil {
		maxLoginCountValue = v.Int()
	}
	return
}

// GetIsReadBodyNameValue Get the data value of the field 'isReadBody' in the gf-security configuration file
func (c *Config) GetIsReadBodyNameValue() (isReadBodyValue bool) {
	v := handlerConfigNotFound(join(configName, isReadBodyName))
	isReadBodyValue = isReadBodyDefValue
	if v != nil {
		isReadBodyValue = v.Bool()
	}
	return
}

// GetIsReadHeadNameValue Get the data value of the field 'isReadHead' in the gf-security configuration file
func (c *Config) GetIsReadHeadNameValue() (isReadHeadValue bool) {
	v := handlerConfigNotFound(join(configName, isReadHeadName))
	isReadHeadValue = isReadHeadDefValue
	if v != nil {
		isReadHeadValue = v.Bool()
	}
	return
}

// GetIsReadCookieNameValue Get the data value of the field 'isReadCookie' in the gf-security configuration file
func (c *Config) GetIsReadCookieNameValue() (isReadCookieValue bool) {
	v := handlerConfigNotFound(join(configName, isReadCookieName))
	isReadCookieValue = isReadCookieDefValue
	if v != nil {
		isReadCookieValue = v.Bool()
	}
	return
}

// GetTokenStyleNameValue Get the data value of the field 'tokenStyle' in the gf-security configuration file
func (c *Config) GetTokenStyleNameValue() (tokenStyleValue string) {
	v := handlerConfigNotFound(join(configName, tokenStyleName))
	tokenStyleValue = tokenStyleDefValue
	if v != nil {
		tokenStyleValue = v.String()
	}
	return
}

// GetDataRefreshPeriodNameValue Get the data value of the field 'dataRefreshPeriod' in the gf-security configuration file
func (c *Config) GetDataRefreshPeriodNameValue() (dataRefreshPeriodValue int) {
	v := handlerConfigNotFound(join(configName, dataRefreshPeriodName))
	dataRefreshPeriodValue = dataRefreshPeriodDefValue
	if v != nil {
		dataRefreshPeriodValue = v.Int()
	}
	return
}

// GetTokenSessionCheckLoginNameValue Get the data value of the field 'tokenSessionCheckLogin' in the gf-security configuration file
func (c *Config) GetTokenSessionCheckLoginNameValue() (tokenSessionCheckLoginValue bool) {
	v := handlerConfigNotFound(join(configName, tokenSessionCheckLoginName))
	tokenSessionCheckLoginValue = tokenSessionCheckLoginDefValue
	if v != nil {
		tokenSessionCheckLoginValue = v.Bool()
	}
	return
}

// GetAutoRenewNameValue Get the data value of the field 'autoRenew' in the gf-security configuration file
func (c *Config) GetAutoRenewNameValue() (autoRenewValue bool) {
	v := handlerConfigNotFound(join(configName, autoRenewName))
	autoRenewValue = autoRenewDefValue
	if v != nil {
		autoRenewValue = v.Bool()
	}
	return
}

// GetTokenPrefixNameValue Get the data value of the field 'tokenPrefix' in the gf-security configuration file
func (c *Config) GetTokenPrefixNameValue() (tokenPrefixValue string) {
	v := handlerConfigNotFound(join(configName, tokenPrefixName))
	tokenPrefixValue = tokenPrefixDefValue
	if v != nil {
		tokenPrefixValue = v.String()
	}
	return
}

// GetIsPrintNameValue Get the data value of the field 'isPrint' in the gf-security configuration file
func (c *Config) GetIsPrintNameValue() (isPrintValue bool) {
	v := handlerConfigNotFound(join(configName, isPrintName))
	isPrintValue = isPrintDefValue
	if v != nil {
		isPrintValue = v.Bool()
	}
	return
}

// GetIsLogNameValue Get the data value of the field 'isLog' in the gf-security configuration file
func (c *Config) GetIsLogNameValue() (isLogValue bool) {
	v := handlerConfigNotFound(join(configName, isLogName))
	isLogValue = isLogDefValue
	if v != nil {
		isLogValue = v.Bool()
	}
	return
}

// GetJwtSecretKeyNameValue Get the data value of the field 'jwtSecretKey' in the gf-security configuration file
func (c *Config) GetJwtSecretKeyNameValue() (jwtSecretKeyValue string) {
	v := handlerConfigNotFound(join(configName, jwtSecretKeyName))
	jwtSecretKeyValue = jwtSecretKeyDefValue
	if v != nil {
		jwtSecretKeyValue = v.String()
	}
	return
}

// GetIdTokenTimeoutNameValue Get the data value of the field 'idTokenTimeout' in the gf-security configuration file
func (c *Config) GetIdTokenTimeoutNameValue() (idTokenTimeoutValue int64) {
	v := handlerConfigNotFound(join(configName, idTokenTimeoutName))
	idTokenTimeoutValue = idTokenTimeoutDefValue
	if v != nil {
		idTokenTimeoutValue = v.Int64()
	}
	return
}

// GetBasicNameValue Get the data value of the field 'basic' in the gf-security configuration file
func (c *Config) GetBasicNameValue() (basicValue string) {
	v := handlerConfigNotFound(join(configName, basicName))
	basicValue = basicDefValue
	if v != nil {
		basicValue = v.String()
	}
	return
}

// GetCurrDomainNameValue Get the data value of the field 'currDomain' in the gf-security configuration file
func (c *Config) GetCurrDomainNameValue() (currDomainValue string) {
	v := handlerConfigNotFound(join(configName, currDomainName))
	currDomainValue = currDomainDefValue
	if v != nil {
		currDomainValue = v.String()
	}
	return
}

// GetCheckIdTokenNameValue Get the data value of the field 'checkIdToken' in the gf-security configuration file
func (c *Config) GetCheckIdTokenNameValue() (checkIdTokenValue bool) {
	v := handlerConfigNotFound(join(configName, checkIdTokenName))
	checkIdTokenValue = checkIdTokenDefValue
	if v != nil {
		checkIdTokenValue = v.Bool()
	}
	return
}

// join Use '.' to concatenate multiple parameters
func join(param ...string) (pattern string) {
	return strings.Join(param, ".")
}

// handlerConfigNotFound handling unable to read configuration file
func handlerConfigNotFound(param string) *gvar.Var {
	v, _ := g.Config(configFileName).Get(SecurityCtx, param)
	if v == nil {
		allField := strings.Split(param, ".")
		length := len(allField)
		printNotConfigured(allField[length-1])
	}
	return v
}

// printNotConfigured output is missing configuration reminder
func printNotConfigured(paramName string) {
	SecurityLog.Warningf(SecurityCtx, "You have not configured '%s', the system will use the default value .", paramName)
}
