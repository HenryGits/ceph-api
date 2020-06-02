package controllers

import (
	"github.com/ceph/go-ceph/system/services"
	"github.com/ceph/go-ceph/system/utils"
	"github.com/ceph/go-ceph/system/web"
	"github.com/kataras/iris"
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

// @tags ceph池模块
// @Summary 获取到所有的pool
// @Description 获取到所有的pool
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /pool/pools [post]
func (c *PoolController) PostPools() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	pools, err := c.PoolService.GetPools(config)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg(pools)
	}
	utils.Log.Info("Response: ", result)
	return result
}

// @tags ceph池模块
// @Summary 通过名称获取对应池信息
// @Description 通过名称获取对应池信息
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /pool/pool [post]
//func (c *PoolController) PostPool(poolName string) *web.ResponseBean {
//	var result *web.ResponseBean
//	//通过context.ReadJSON()读取传过来的数据
//	var config web.ConnConfig
//	if err := c.Ctx.ReadJSON(&config); err != nil {
//		utils.Log.Error(err)
//	}
//	pools, err := c.PoolService.GetPoolByName(config, poolName)
//	if err != nil {
//		result = web.GenFailedMsg(err.Error())
//	} else {
//		result = web.GenSuccessMsg(pools)
//	}
//	utils.Log.Info("Response: ", result)
//	return result
//}

// @tags ceph池模块
// @Summary 创建池
// @Description 创建池
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /pool/pool [post]
func (c *PoolController) PostPool() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	err := c.PoolService.AddPool(config, poolName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccess("池创建成功。")
	}
	utils.Log.Info("Response: ", result)
	return result
}

// @tags ceph池模块
// @Summary 删除池
// @Description 删除池
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /pool/pool [delete]
func (c *PoolController) DeletePool() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	err := c.PoolService.DeletePoolByName(config, poolName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg("池删除成功。")
	}
	utils.Log.Info("Response: ", result)
	return result
}
