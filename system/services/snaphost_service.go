package services

import (
	"ceph-api/system/utils"
	"ceph-api/system/web"
	"errors"
	"github.com/ceph/go-ceph/rbd"
)

/*
	Ceph Rbd Image SnaphostApi操作
*/
type ISnaphostService interface {
	// 获取所有的rbd image快照
	GetSnaphosts(config web.ConnConfig, poolName, imageName string) ([]rbd.SnapInfo, error)

	// 获取rbd image快照
	GetSnaphost(config web.ConnConfig, poolName, imageName, snapName string) (*rbd.Snapshot, error)

	//创建rbd image快照
	CreateSnaphost(config web.ConnConfig, poolName, imageName, snapName string) (*rbd.Snapshot, error)

	//删除rbd image快照
	DeleteSnaphost(config web.ConnConfig, poolName, imageName, snapName string) error

	//锁定快照
	ProtectSnaphost(config web.ConnConfig, poolName, imageName, snapName string) error

	// 解除快照锁定
	UnProtectSnaphost(config web.ConnConfig, poolName, imageName, snapName string) error

	// 快照回滚
	Rollback(config web.ConnConfig, poolName, imageName, snapName string) error

	//将快照克隆生成新的rbd image
	CloneSnaphost(config web.ConnConfig, poolName, oldImageName, newImageName, snapName string) (*rbd.Image, error)
}

// NewPoolService returns the default pool service.
func NewSnaphostService() ISnaphostService {
	return &snaphostService{}
}

type snaphostService struct {
}

func (s snaphostService) GetSnaphosts(config web.ConnConfig, poolName, imageName string) ([]rbd.SnapInfo, error) {
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
	img, err := rbd.OpenImage(ioctx, imageName, "")
	if err != nil {
		utils.Log.Error("打开Rbd image:「" + imageName + "」异常: " + err.Error())
		return nil, errors.New("打开打开Rbd image异常！")
	}
	snaps, err := img.GetSnapshotNames()
	return snaps, err
}

func (s snaphostService) GetSnaphost(config web.ConnConfig, poolName, imageName, snapName string) (*rbd.Snapshot, error) {
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
	img, err := rbd.OpenImage(ioctx, imageName, "")
	if err != nil {
		utils.Log.Error("打开Rbd image:「" + imageName + "」异常: " + err.Error())
		return nil, errors.New("打开打开Rbd image异常！")
	}
	snap := img.GetSnapshot(snapName)
	return snap, err
}

func (s snaphostService) CreateSnaphost(config web.ConnConfig, poolName, imageName, snapName string) (*rbd.Snapshot, error) {
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
	img, err := rbd.OpenImage(ioctx, imageName, "")
	if err != nil {
		utils.Log.Error("打开Rbd image:「" + imageName + "」异常: " + err.Error())
		return nil, errors.New("打开打开Rbd image异常！")
	}
	return img.CreateSnapshot(snapName)
}

func (s snaphostService) DeleteSnaphost(config web.ConnConfig, poolName, imageName, snapName string) error {
	conn, err := utils.GetConnection(config)
	if err != nil {
		utils.Log.Error("获取连接异常: ", err)
		return err
	}
	// connect to the pool
	ioctx, err := conn.OpenIOContext(poolName)
	if err != nil {
		utils.Log.Error("打开pool: {}异常: ", poolName, err)
		return err
	}
	img, err := rbd.OpenImage(ioctx, imageName, "")
	if err != nil {
		utils.Log.Error("打开Rbd image:「" + imageName + "」异常: " + err.Error())
		return errors.New("打开打开Rbd image异常！")
	}
	snap := img.GetSnapshot(snapName)
	if flag, _ := snap.IsProtected(); flag {
		return errors.New("快照被锁定,无法删除！")
	}
	return snap.Remove()
}

func (s snaphostService) ProtectSnaphost(config web.ConnConfig, poolName, imageName, snapName string) error {
	conn, err := utils.GetConnection(config)
	if err != nil {
		utils.Log.Error("获取连接异常: ", err)
		return err
	}
	// connect to the pool
	ioctx, err := conn.OpenIOContext(poolName)
	if err != nil {
		utils.Log.Error("打开pool: {}异常: ", poolName, err)
		return err
	}
	img, err := rbd.OpenImage(ioctx, imageName, "")
	if err != nil {
		utils.Log.Error("打开Rbd image:「" + imageName + "」异常: " + err.Error())
		return errors.New("打开打开Rbd image异常！")
	}
	snap := img.GetSnapshot(snapName)
	return snap.Protect()
}

func (s snaphostService) UnProtectSnaphost(config web.ConnConfig, poolName, imageName, snapName string) error {
	conn, err := utils.GetConnection(config)
	if err != nil {
		utils.Log.Error("获取连接异常: ", err)
		return err
	}
	// connect to the pool
	ioctx, err := conn.OpenIOContext(poolName)
	if err != nil {
		utils.Log.Error("打开pool: {}异常: ", poolName, err)
		return err
	}
	img, err := rbd.OpenImage(ioctx, imageName, "")
	if err != nil {
		utils.Log.Error("打开Rbd image:「" + imageName + "」异常: " + err.Error())
		return errors.New("打开打开Rbd image异常！")
	}
	snap := img.GetSnapshot(snapName)
	return snap.Unprotect()
}

func (s snaphostService) Rollback(config web.ConnConfig, poolName, imageName, snapName string) error {
	conn, err := utils.GetConnection(config)
	if err != nil {
		utils.Log.Error("获取连接异常: ", err)
		return err
	}
	// connect to the pool
	ioctx, err := conn.OpenIOContext(poolName)
	if err != nil {
		utils.Log.Error("打开pool: {}异常: ", poolName, err)
		return err
	}
	img, err := rbd.OpenImage(ioctx, imageName, "")
	if err != nil {
		utils.Log.Error("打开Rbd image:「" + imageName + "」异常: " + err.Error())
		return errors.New("打开Rbd image异常！")
	}
	snap := img.GetSnapshot(snapName)
	go snap.Rollback()
	return err
}

func (s snaphostService) CloneSnaphost(config web.ConnConfig, poolName, oldImageName, newImageName, snapName string) (*rbd.Image, error) {
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
	image, err := rbd.OpenImage(ioctx, oldImageName, "")
	if err != nil {
		utils.Log.Error("打开Rbd image:「" + oldImageName + "」异常: " + err.Error())
		return nil, errors.New("打开Rbd image异常！")
	}
	snap := image.GetSnapshot(snapName)
	// 锁定 snapshot
	if err := snap.Protect(); err != nil {
		utils.Log.Error("快照锁定失败！", err.Error())
		return nil, err
	}
	return image.Clone(snapName, ioctx, newImageName, rbd.RbdFeatureLayering, 22)
}
