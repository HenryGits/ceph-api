package services

import (
	"github.com/ceph/go-ceph/system/web"
	"github.com/kataras/iris"
)

/**
 * 日志模块
 */
var rbdLog = iris.New().Logger()

/*
	Ceph Rbd Api操作
*/
type IRbdService interface {
	// 获取所有的rbd image
	GetImages(config web.ConnConfig) ([]string, error)

	// 获取所有的rbd image
	GetImageByName(config web.ConnConfig, poolName string) (int64, error)

	//创建rbd image
	CreateRbdImage(poolName, imageName string) error

	//删除rbd image
	DeleteImageByName(poolName, imageName string) error
}

// NewPoolService returns the default pool service.
func NewRbdService() IRbdService {
	return &rbdService{}
}

type rbdService struct {
}

func (r rbdService) GetImages(config web.ConnConfig) ([]string, error) {
	panic("implement me")
}

func (r rbdService) GetImageByName(config web.ConnConfig, poolName string) (int64, error) {
	panic("implement me")
}

func (r rbdService) CreateRbdImage(poolName, imageName string) error {
	panic("implement me")
}

func (r rbdService) DeleteImageByName(poolName, imageName string) error {
	panic("implement me")
}
