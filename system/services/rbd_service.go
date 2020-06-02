package services

import (
	"github.com/ceph/go-ceph/rbd"
	"github.com/ceph/go-ceph/system/utils"
	"github.com/ceph/go-ceph/system/web"
)

/*
	Ceph Rbd Api操作
*/
type IRbdService interface {
	// 获取所有的rbd image
	GetImages(config web.ConnConfig, poolName string) ([]string, error)

	// 获取所有的rbd image
	GetImageByName(config web.ConnConfig, poolName string) *rbd.Image

	//创建rbd image
	CreateRbdImage(config web.ConnConfig, poolName, imageName string) (*rbd.Image, error)

	//删除rbd image
	DeleteImageByName(config web.ConnConfig, poolName, imageName string) error
}

// NewPoolService returns the default pool service.
func NewRbdService() IRbdService {
	return &rbdService{}
}

type rbdService struct {
}

func (r rbdService) GetImages(config web.ConnConfig, poolName string) ([]string, error) {
	conn, err := utils.GetConnection(config)
	if err != nil {
		utils.Log.Error("获取连接异常: ", err)
		return nil, err
	}
	// connect to the pool
	ioctx, err := conn.OpenIOContext(poolName)
	if err != nil {
		utils.Log.Error("打开pool: {}异常: ", poolName, err)
		return nil, err
	}
	return rbd.GetImageNames(ioctx)
}

func (r rbdService) GetImageByName(config web.ConnConfig, poolName string) *rbd.Image {
	conn, err := utils.GetConnection(config)
	if err != nil {
		utils.Log.Error("获取连接异常: ", err)
		return nil
	}
	// connect to the pool
	ioctx, err := conn.OpenIOContext(poolName)
	if err != nil {
		utils.Log.Error("打开pool: {}异常: ", poolName, err)
		return nil
	}
	return rbd.GetImage(ioctx, poolName)
}

func (r rbdService) CreateRbdImage(config web.ConnConfig, poolName, imageName string) (*rbd.Image, error) {
	conn, err := utils.GetConnection(config)
	if err != nil {
		utils.Log.Error("获取连接异常: ", err)
		return nil, err
	}
	// connect to the pool
	ioctx, err := conn.OpenIOContext(poolName)
	if err != nil {
		utils.Log.Error("打开pool: {}异常: ", poolName, err)
		return nil, err
	}
	return rbd.Create2(ioctx, imageName, utils.DefaultBaseImageSize, rbd.FeatureLayering, 22)
}

func (r rbdService) DeleteImageByName(config web.ConnConfig, poolName, imageName string) error {
	conn, err := utils.GetConnection(config)
	if err != nil {
		utils.Log.Error("获取连接异常: ", err)
		return nil
	}
	// connect to the pool
	ioctx, err := conn.OpenIOContext(poolName)
	if err != nil {
		utils.Log.Error("打开pool: {}异常: ", poolName, err)
		return nil
	}
	return rbd.RemoveImage(ioctx, poolName)
}
