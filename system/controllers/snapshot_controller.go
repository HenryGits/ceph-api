package controllers

import (
	"github.com/ceph/go-ceph/system/services"
	"github.com/ceph/go-ceph/system/web"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type SnaphostController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding)
	// and the Session which depends on the current context (dynamic binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	SnaphostService services.ISnaphostService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

/**
 * 日志模块
 */
var snapLog = iris.New().Logger()

// @tags ceph snaphost模块
// @Summary 获取到所有的快照
// @Description 获取到所有的快照
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName path string true "池名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /snap/{poolName}/{imageName}/Snaphosts [get]
func (c *SnaphostController) GetSnaphosts(config web.ConnConfig, poolName, imageName string) *web.ResponseBean {
	var result *web.ResponseBean
	images, err := c.SnaphostService.GetSnaphosts(config, poolName, imageName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg(images)
	}
	log.Info("Response: ", result)
	return result
}

// @tags ceph snaphost模块
// @Summary 获取快照信息
// @Description 获取快照信息
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName path string true "池名称"
// @Param poolName path string true "image名称"
// @Param snapName path string true "快照名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /rbd/{poolName}/{imageName}/{snapName} [get]
func (c *SnaphostController) GetSnaphost(config web.ConnConfig, poolName, imageName, snapName string) *web.ResponseBean {
	var result *web.ResponseBean
	pools, err := c.SnaphostService.GetSnaphost(config, poolName, imageName, snapName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg(pools)
	}
	log.Info("Response: ", result)
	return result
}

// @tags ceph snaphost模块
// @Summary 创建快照
// @Description 创建快照
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName path string true "池名称"
// @Param poolName path string true "image名称"
// @Param snapName path string true "快照名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /rbd/{poolName}/{imageName}/{snapName} [post]
func (c *SnaphostController) CreateSnaphost(config web.ConnConfig, poolName, imageName, snapName string) *web.ResponseBean {
	var result *web.ResponseBean
	err := c.SnaphostService.CreateSnaphost(config, poolName, imageName, snapName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccess("池创建成功。")
	}
	log.Info("Response: ", result)
	return result
}

// @tags ceph snaphost模块
// @Summary 删除快照
// @Description 删除快照
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName path string true "池名称"
// @Param poolName path string true "image名称"
// @Param snapName path string true "快照名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /rbd/{poolName}/{imageName}/{snapName} [delete]
func (c *SnaphostController) DeleteSnaphost(config web.ConnConfig, poolName, imageName, snapName string) *web.ResponseBean {
	var result *web.ResponseBean
	err := c.SnaphostService.DeleteSnaphost(config, poolName, imageName, snapName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg("删除池成功。")
	}
	log.Info("Response: ", result)
	return result
}
