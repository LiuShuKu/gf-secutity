// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
// Project Name :   gf-security
// Author       :   liushuku@yeah.net
package gresult_test

import (
	"context"
	"fmt"
	"gf-secutity/gresult"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
	"testing"
)

var (
	ctx  = context.TODO()
	port = 9927
	log  = g.Log()
)

// struct for testing --> TestResponse
type user struct {
	Name     string `json:"name"`
	NickName string `json:"nickName"`
}

// struct for testing --> TestResponseTotal
type student struct {
	StudentName string `json:"studentName"`
	StudentAge  int    `json:"studentAge"`
}

func TestSuccess(t *testing.T) {

	s := g.Server(guid.S())
	{
		s.BindHandler("/", func(r *ghttp.Request) {
			gresult.Success(r)
		})
		s.SetPort(port)
	}

	{
		_ = s.Start()
		defer func() {
			_ = s.Shutdown()
		}()
	}

	{
		gtest.C(t, func(t *gtest.T) {
			c := g.Client()
			log.Info(ctx, c.GetContent(ctx, "http://127.0.0.1:9927"))
			// print : {"code":0,"msg":"OK","data":null}
		})
	}

}
func TestSuccessData(t *testing.T) {

	data := user{
		Name:     "GuShuLiu",
		NickName: "LiuShuKu",
	}

	s := g.Server(guid.S())
	{
		s.BindHandler("/", func(r *ghttp.Request) {
			gresult.SuccessData(data, r)
		})
		s.SetPort(port)
	}

	{
		_ = s.Start()
		defer func() {
			_ = s.Shutdown()
		}()
	}

	{
		gtest.C(t, func(t *gtest.T) {
			c := g.Client()
			log.Info(ctx, c.GetContent(ctx, "http://127.0.0.1:9927"))
			// print : {"code":0,"msg":"OK","data":{"name":"GuShuLiu","nickName":"LiuShuKu"}}
		})
	}

}

func TestResponseTotal(t *testing.T) {
	stu0 := student{
		StudentName: "ZhangSan",
		StudentAge:  18,
	}
	stu1 := student{
		StudentName: "GuShuLiu",
		StudentAge:  19,
	}
	students := []student{stu0, stu1}

	s := g.Server(guid.S())
	{
		s.BindHandler("/", func(r *ghttp.Request) {
			gresult.SuccessDataTotal(students, 2, r)
		})

		s.SetPort(port)
	}

	{
		_ = s.Start()
		defer func() {
			_ = s.Shutdown()
		}()
	}
	{
		gtest.C(t, func(t *gtest.T) {
			c := g.Client()
			fmt.Println(c.PostContent(ctx, "http://127.0.0.1:9927"))
			// print : {"code":0,"msg":"OK","data":[{"studentName":"ZhangSan","studentAge":18},{"studentName":"GuShuLiu","studentAge":19}],"total":2}
		})
	}
}

func TestSuccessDataTotalInner(t *testing.T) {
	stu0 := student{
		StudentName: "ZhangSan",
		StudentAge:  18,
	}
	stu1 := student{
		StudentName: "GuShuLiu",
		StudentAge:  19,
	}
	students := []student{stu0, stu1}

	s := g.Server(guid.S())
	{
		s.BindHandler("/", func(r *ghttp.Request) {
			gresult.SuccessDataTotalInner(students, 2, r)
		})

		s.SetPort(port)
	}

	{
		_ = s.Start()
		defer func() {
			_ = s.Shutdown()
		}()
	}
	{
		gtest.C(t, func(t *gtest.T) {
			c := g.Client()
			fmt.Println(c.PostContent(ctx, "http://127.0.0.1:9927"))
			// {"code":0,"msg":"OK","data":{"rows":[{"studentName":"ZhangSan","studentAge":18},{"studentName":"GuShuLiu","studentAge":19}],"total":2}}
		})
	}
}
