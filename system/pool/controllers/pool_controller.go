// file: controllers/pool_controller.go

package controllers

import (
	"github.com/ceph/go-ceph/system/pool/services"
	_ "github.com/ceph/go-ceph/system/web"
	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type PoolController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding)
	// and the Session which depends on the current context (dynamic binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	PoolService services.IPoolService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

const userIDKey = "UserID"

/**
 * 日志模块
 */
var log = iris.New().Logger()

func (c *PoolController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(userIDKey, 0)
	return userID
}

func (c *PoolController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

func (c *PoolController) logout() {
	c.Session.Destroy()
}

// GetPools handles GET: http://localhost:8080/v1/pool/pools.
// @tags ceph池模块
// @Summary 获取到所有的pool
// @Description 获取到所有的pool
// @Accept  application/json
// @Produce  application/json
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Failure 500 {object} web.APIError "Server Error"
// @Router /pool/pools [get]
func (c *PoolController) GetPools() mvc.Result {
	var result mvc.Result
	pools, err := c.PoolService.GetPools()
	if err != nil {
		result = mvc.View{
			Code: 500,
			Err:  err,
		}
	} else {
		jsonStr, _ := jsoniter.Marshal(pools)
		result = mvc.View{
			Code: 200,
			Data: jsonStr,
		}
	}
	log.Info("Response: ", result)
	return result
}
