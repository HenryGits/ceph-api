package services

import (
	"github.com/ceph/go-ceph/rbd"
	"github.com/ceph/go-ceph/system/web"
	"github.com/kataras/iris"
)

/**
 * 日志模块
 */
var snapLog = iris.New().Logger()

/*
	Ceph Rbd Api操作
*/
type ISnaphostService interface {
	// 获取所有的rbd image快照
	GetSnaphosts(config web.ConnConfig, poolName, imageName string) ([]string, error)

	// 获取rbd image快照
	GetSnaphost(config web.ConnConfig, poolName, imageName, snapName string) (rbd.Snapshot, error)

	//创建rbd image快照
	CreateSnaphost(config web.ConnConfig, poolName, imageName, snapName string) error

	//删除rbd image快照
	DeleteSnaphost(config web.ConnConfig, poolName, imageName, snapName string) error

	//锁定快照
	ProtectSnaphost(config web.ConnConfig, poolName, imageName, newImageName, snapName string) error

	//将快照克隆生成rbd image
	CloneSnaphost(config web.ConnConfig, poolName, oldImageName, newImageName, snapName string) error
}

// NewPoolService returns the default pool service.
func NewSnaphostService() ISnaphostService {
	return &snaphostService{}
}

type snaphostService struct {
}

func (s snaphostService) GetSnaphosts(config web.ConnConfig, poolName, imageName string) ([]string, error) {
	panic("implement me")
}

func (s snaphostService) GetSnaphost(config web.ConnConfig, poolName, imageName, snapName string) (rbd.Snapshot, error) {
	panic("implement me")
}

func (s snaphostService) CreateSnaphost(config web.ConnConfig, poolName, imageName, snapName string) error {
	panic("implement me")
}

func (s snaphostService) DeleteSnaphost(config web.ConnConfig, poolName, imageName, snapName string) error {
	panic("implement me")
}

func (s snaphostService) ProtectSnaphost(config web.ConnConfig, poolName, imageName, newImageName, snapName string) error {
	panic("implement me")
}

func (s snaphostService) CloneSnaphost(config web.ConnConfig, poolName, oldImageName, newImageName, snapName string) error {
	panic("implement me")
}
