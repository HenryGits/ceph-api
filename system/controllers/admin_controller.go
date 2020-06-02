package controllers

import (
	"github.com/ceph/go-ceph/system/web"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"time"
)

type AdminController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding)
	// and the Session which depends on the current context (dynamic binding).
	Ctx iris.Context

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

// Auth godoc
// @Summary Auth admin
// @Description get admin info
// @Tags accounts,admin
// @Accept  json
// @Produce  json
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 401 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Security ApiKeyAuth
// @Router /admin/auth [post]
func (c *AdminController) Auth() *web.ResponseBean {
	phone := c.Ctx.FormValue("phone")
	password := c.Ctx.FormValue("password")

	/* 这里省去了对用户的验证，在实际使用过程中需要验证用户是否存在，密码是否正确 */
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone":    phone,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // 添加过期时间为2个小时
	})

	// 这里的密钥和前面的必须一样
	tokenString, _ := token.SignedString([]byte("My Secret"))
	return web.GenSuccessMsg(tokenString)
}
