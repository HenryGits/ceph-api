package controllers

import (
	"github.com/ceph/go-ceph/system/services"
	"github.com/ceph/go-ceph/system/utils"
	"github.com/ceph/go-ceph/system/web"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
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

// @tags ceph snaphost模块
// @Summary 获取到所有的快照
// @Description 获取到所有的快照
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Param imageName query string true "image名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /snap/snaphosts [post]
func (c *SnaphostController) PostSnaphosts() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	var imageName = c.Ctx.URLParam("imageName")
	images, err := c.SnaphostService.GetSnaphosts(config, poolName, imageName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg(images)
	}
	utils.Log.Info("Response: ", result)
	return result
}

// @tags ceph snaphost模块
// @Summary 获取快照信息
// @Description 获取快照信息
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Param imageName query string true "image名称"
// @Param snapName query string true "快照名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /snap/snaphost [post]
func (c *SnaphostController) PostSnaphost() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	var imageName = c.Ctx.URLParam("imageName")
	var snapName = c.Ctx.URLParam("snapName")
	snap, err := c.SnaphostService.GetSnaphost(config, poolName, imageName, snapName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg(snap)
	}
	utils.Log.Info("Response: ", result)
	return result
}

// @tags ceph snaphost模块
// @Summary 创建快照
// @Description 创建快照
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Param imageName query string true "image名称"
// @Param snapName query string true "快照名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /snap/snaphost/create [post]
func (c *SnaphostController) PostSnaphostCreate() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	var imageName = c.Ctx.URLParam("imageName")
	var snapName = c.Ctx.URLParam("snapName")
	snap, err := c.SnaphostService.CreateSnaphost(config, poolName, imageName, snapName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg(snap)
	}
	utils.Log.Info("Response: ", result)
	return result
}

// @tags ceph snaphost模块
// @Summary 删除快照
// @Description 删除快照
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Param imageName query string true "image名称"
// @Param snapName query string true "快照名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /snap/snaphost [delete]
func (c *SnaphostController) DeleteSnaphost() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	var imageName = c.Ctx.URLParam("imageName")
	var snapName = c.Ctx.URLParam("snapName")
	err := c.SnaphostService.DeleteSnaphost(config, poolName, imageName, snapName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg("快照删除成功。")
	}
	utils.Log.Info("Response: ", result)
	return result
}

// @tags ceph snaphost模块
// @Summary 锁定快照
// @Description 锁定快照
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Param imageName query string true "image名称"
// @Param snapName query string true "快照名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /snap/protect [post]
func (c *SnaphostController) PostProtect() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	var imageName = c.Ctx.URLParam("imageName")
	var snapName = c.Ctx.URLParam("snapName")
	err := c.SnaphostService.ProtectSnaphost(config, poolName, imageName, snapName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg("快照锁定成功。")
	}
	utils.Log.Info("Response: ", result)
	return result
}

// @tags ceph snaphost模块
// @Summary 解锁快照
// @Description 解锁快照
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Param imageName query string true "image名称"
// @Param snapName query string true "快照名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /snap/unprotect [post]
func (c *SnaphostController) PostUnprotect() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	var imageName = c.Ctx.URLParam("imageName")
	var snapName = c.Ctx.URLParam("snapName")
	err := c.SnaphostService.UnProtectSnaphost(config, poolName, imageName, snapName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccess("快照解锁成功。")
	}
	utils.Log.Info("Response: ", result)
	return result
}

// @tags ceph snaphost模块
// @Summary 快照回滚
// @Description 快照回滚
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Param imageName query string true "image名称"
// @Param snapName query string true "快照名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /snap/rollback [post]
func (c *SnaphostController) PostRollback() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	var imageName = c.Ctx.URLParam("imageName")
	var snapName = c.Ctx.URLParam("snapName")
	err := c.SnaphostService.Rollback(config, poolName, imageName, snapName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg("快照回滚成功。")
	}
	utils.Log.Info("Response: ", result)
	return result
}

// @tags ceph snaphost模块
// @Summary 快照克隆
// @Description 快照克隆
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Param oldImageName query string true "image名称"
// @Param newImageName query string true "新的image名称"
// @Param snapName query string true "快照名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /snap/clone [post]
func (c *SnaphostController) PostClone() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	var oldImageName = c.Ctx.URLParam("oldImageName")
	var newImageName = c.Ctx.URLParam("newImageName")
	var snapName = c.Ctx.URLParam("snapName")
	image, err := c.SnaphostService.CloneSnaphost(config, poolName, oldImageName, newImageName, snapName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg(image)
	}
	utils.Log.Info("Response: ", result)
	return result
}
