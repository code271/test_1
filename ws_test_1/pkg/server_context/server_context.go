package server_context

import (
	"code271/ws_test_1/pkg/jwt_helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Context 自定义gin请求上下文，重写方法
type Context struct {
	*gin.Context
	UserID      int64  // 用户id
	AccountName string // 用户名
	RoleID      int32  // 用户角色权限
}

// JSON json输出
func (c *Context) JSON(obj *Response) {
	if obj == nil {
		panic("响应数据对象不能未nil")
	}
	status := http.StatusOK
	c.Context.JSON(status, obj)
}

type HandlerFunc func(c *Context)

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			Context: c,
		}
		// 用户信息
		userinfoObj, ext := c.Get("loginUser")
		if ext == true {
			userinfo := userinfoObj.(*jwt_helper.CustomClaims)
			ctx.UserID = userinfo.UserID
			ctx.AccountName = userinfo.AccountName
			ctx.RoleID = userinfo.RoleID
		}
		h(ctx)
	}
}


