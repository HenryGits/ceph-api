package controllers

import (
	"github.com/ceph/go-ceph/system/services"
	"github.com/ceph/go-ceph/system/web"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type RbdController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding)
	// and the Session which depends on the current context (dynamic binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	RbdService services.IRbdService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

/**
 * 日志模块
 */
var rbdLog = iris.New().Logger()

// @tags ceph rbd模块
// @Summary 获取到所有的image
// @Description 获取到所有的image
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName path string true "池名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /rbd/images [get]
func (c *RbdController) GetImages(config web.ConnConfig, poolName string) *web.ResponseBean {
	var result *web.ResponseBean
	images, err := c.RbdService.GetImages(config)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg(images)
	}
	log.Info("Response: ", result)
	return result
}

// @tags ceph rbd模块
// @Summary 获取image信息
// @Description 获取image信息
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName path string true "池名称"
// @Param poolName path string true "image名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /rbd/{poolName}/{imageName} [get]
func (c *RbdController) GetImageByName(config web.ConnConfig, poolName, imageName string) *web.ResponseBean {
	var result *web.ResponseBean
	pools, err := c.RbdService.GetImageByName(config, poolName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg(pools)
	}
	log.Info("Response: ", result)
	return result
}

// @tags ceph rbd模块
// @Summary 创建image
// @Description 创建image
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName path string true "池名称"
// @Param poolName path string true "image名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /rbd/{poolName}/{imageName} [post]
func (c *RbdController) CreateImage(config web.ConnConfig, poolName, imageName string) *web.ResponseBean {
	var result *web.ResponseBean
	err := c.RbdService.CreateRbdImage(poolName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccess("池创建成功。")
	}
	log.Info("Response: ", result)
	return result
}

// @tags ceph rbd模块
// @Summary 删除池
// @Description 删除池
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName path string true "池名称"
// @Param poolName path string true "image名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /rbd/{poolName}/{imageName} [delete]
func (c *RbdController) DeleteImage(config web.ConnConfig, poolName, imageName string) *web.ResponseBean {
	var result *web.ResponseBean
	err := c.RbdService.DeleteImageByName(poolName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg("删除池成功。")
	}
	log.Info("Response: ", result)
	return result
}
