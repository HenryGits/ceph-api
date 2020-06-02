package controllers

import (
	"github.com/ceph/go-ceph/system/services"
	"github.com/ceph/go-ceph/system/utils"
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

// @tags ceph rbd模块
// @Summary 获取到所有的image
// @Description 获取到所有的image
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /rbd/images [post]
func (c *RbdController) PostImages() *web.ResponseBean {
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	utils.Log.Info(config)
	utils.Log.Info(poolName)

	var result *web.ResponseBean
	images, err := c.RbdService.GetImages(config, poolName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg(images)
	}
	utils.Log.Info("Response: ", result)
	return result
}

// @tags ceph rbd模块
// @Summary 创建image
// @Description 创建image
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Param imageName query string true "image名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /rbd/image [post]
func (c *RbdController) PostImage() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	var imageName = c.Ctx.URLParam("imageName")
	image, err := c.RbdService.CreateRbdImage(config, poolName, imageName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg(image)
	}
	utils.Log.Info("Response: ", result)
	return result
}

// @tags ceph rbd模块
// @Summary 删除image
// @Description 删除image
// @Accept  application/json
// @Produce  application/json
// @Param config body web.ConnConfig true "连接配置"
// @Param poolName query string true "池名称"
// @Param imageName query string true "image名称"
// @Success 200 {object} web.ResponseBean
// @Failure 400 {object} web.ResponseBean
// @Failure 404 {object} web.ResponseBean
// @Failure 500 {object} web.ResponseBean
// @Router /rbd/image [delete]
func (c *RbdController) DeleteImage() *web.ResponseBean {
	var result *web.ResponseBean
	//通过context.ReadJSON()读取传过来的数据
	var config web.ConnConfig
	if err := c.Ctx.ReadJSON(&config); err != nil {
		utils.Log.Error(err)
	}
	var poolName = c.Ctx.URLParam("poolName")
	var imageName = c.Ctx.URLParam("imageName")
	err := c.RbdService.DeleteImageByName(config, poolName, imageName)
	if err != nil {
		result = web.GenFailedMsg(err.Error())
	} else {
		result = web.GenSuccessMsg("删除image成功。")
	}
	utils.Log.Info("Response: ", result)
	return result
}
