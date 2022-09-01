// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.

// Copyright (c)  2022 LiuShuKu
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/LiuShuKu/gf-secutity.
// Project Name :   gf-security
// Author       :   liushuku@yeah.net

package gweb_test

import (
	"context"
	"gf-secutity/gresult"
	"gf-secutity/gweb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
	"net/http"
	"testing"
)

var (
	ctx  = context.TODO()
	port = 9927
	log  = g.Log()
)

func TestHandlerStatus400(t *testing.T) {

	s := g.Server(guid.S())
	{

		s.Use(gweb.HandlerStatus400)
		s.SetPort(port)
		s.BindHandler("/", func(r *ghttp.Request) {
			// manually set the status code
			r.Response.Status = http.StatusBadRequest
		})
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
			log.Info(ctx, c.GetContent(ctx, "http://127.0.0.1:9927/"))
			// print: {"code":2,"msg":"Bad Request","data":null}
		})
	}
}
func TestHandlerStatus401(t *testing.T) {

	s := g.Server(guid.S())
	{
		s.Use(gweb.HandlerStatus401)
		s.SetPort(port)
		s.BindHandler("/", func(r *ghttp.Request) {
			// manually set the status code
			r.Response.Status = http.StatusUnauthorized
		})
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
			log.Info(ctx, c.GetContent(ctx, "http://127.0.0.1:9927/"))
			// print: {"code":-1,"msg":"Unauthorized","data":null}
		})
	}
}
func TestHandlerStatus405(t *testing.T) {

	s := g.Server(guid.S())
	{
		s.Use(gweb.HandlerStatus405)
		s.SetPort(port)

		s.BindHandler("/", func(r *ghttp.Request) {
			// manually set the status code
			r.Response.Status = http.StatusMethodNotAllowed
		})
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
			log.Info(ctx, c.GetContent(ctx, "http://127.0.0.1:9927/"))
			// print: {"code":2,"msg":"Method Not Allowed","data":null}
		})
	}
}
func TestHandlerStatus415(t *testing.T) {

	s := g.Server(guid.S())
	{
		s.Use(gweb.HandlerStatus415)
		s.SetPort(port)

		s.BindHandler("/", func(r *ghttp.Request) {
			// manually set the status code
			r.Response.Status = http.StatusUnsupportedMediaType
		})
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
			log.Info(ctx, c.GetContent(ctx, "http://127.0.0.1:9927/"))
			// print: {"code":2,"msg":"Unsupported Media Type","data":null}
		})
	}
}
func TestHandlerStatus404(t *testing.T) {

	s := g.Server(guid.S())
	{
		s.BindHandler("/", func(r *ghttp.Request) {
			gresult.Success(r)
		})
		s.Use(gweb.HandlerStatus404)
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
			//  path "/xxx" does not exist and a 404 will be thrown here
			log.Info(ctx, c.GetContent(ctx, "http://127.0.0.1:9927/xxx"))
			// print: {"code":2,"msg":"Not Found","data":null}
		})
	}
}
func TestHandlerStatus500(t *testing.T) {

	s := g.Server(guid.S())
	{
		s.BindHandler("/", func(r *ghttp.Request) {
			// simulate error here system error
			// note: the system panic will exit the program!!!
			panic("simulation error! ")
		})
		// This middleware should be used in a production environment
		// it doesn't expose what s in the error message
		s.Use(gweb.HandlerStatus500)

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
			log.Info(ctx, c.GetContent(ctx, "http://127.0.0.1:9927/"))
			// print: {"code":2,"msg":"Internal Server Error","data":null}
		})
	}
}
func TestHandlerServer500(t *testing.T) {

	s := g.Server(guid.S())
	{
		s.BindHandler("/", func(r *ghttp.Request) {
			// simulate error here system error
			// note: the system panic will exit the program!!!
			panic("simulation error! ")
		})

		//  The middleware will throw an error message: "simulation error! " will be printed
		s.Use(gweb.HandlerServer500)

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
			log.Info(ctx, c.GetContent(ctx, "http://127.0.0.1:9927/"))
			// print: {"code":2,"msg":"exception recovered: simulation error! ","data":null}
		})
	}
}
