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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	SecurityLog = logConfig()
)

// logConfig : just added the prefix
func logConfig() (logger *glog.Logger) {
	logger = g.Log(configName)
	{
		config := logger.GetConfig()
		config.Prefix = logPrefix
		_ = logger.SetConfig(config)
	}
	return
}
