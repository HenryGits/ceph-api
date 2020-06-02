package services

import (
	"github.com/ceph/go-ceph/system/utils"
	"github.com/ceph/go-ceph/system/web"
	"github.com/kataras/iris"
)

/*
	Ceph Pool Api操作
*/
type IPoolService interface {
	/*
		/api/pool
	*/
	GetPools(config web.ConnConfig) ([]string, error)

	/*
		/api/pool/{pool_name}
	*/
	GetPoolByName(config web.ConnConfig, poolName string) (int64, error)

	/*
		创建ceph存储池
		/api/pool
	*/
	AddPool(config web.ConnConfig, poolName string) error

	/*
		通过poolName删除数据库
	*/
	DeletePoolByName(config web.ConnConfig, poolName string) error
}

// NewPoolService returns the default pool service.
func NewPoolService() IPoolService {
	return &poolService{}
}

type poolService struct {
}

func (p poolService) GetPools(config web.ConnConfig) ([]string, error) {
	utils.Log.Info("==>GetPools")

	//获取调go-ceph获取pool
	conn, err := utils.GetConnection(config)
	utils.Log.Info("获取连接成功! GetInstanceID: ", conn.GetInstanceID())
	if err != nil {
		utils.Log.Error("获取连接异常: ", err)
		return nil, err
	}
	pools, err := conn.ListPools()
	utils.Log.Info("Pools: ", pools)
	return pools, err
}

func (p poolService) GetPoolByName(config web.ConnConfig, poolName string) (int64, error) {
	conn, err := utils.GetConnection(config)
	if err != nil {
		utils.Log.Error("获取连接异常: ", err)
		return 0, err
	}
	return conn.GetPoolByName(poolName)
}

func (p poolService) AddPool(config web.ConnConfig, poolName string) error {
	conn, err := utils.GetConnection(config)
	if err != nil {
		utils.Log.Error("获取连接异常: ", err)
		return err
	}
	return conn.MakePool(poolName)
}

func (p poolService) DeletePoolByName(config web.ConnConfig, poolName string) error {
	conn, err := utils.GetConnection(config)
	if err != nil {
		utils.Log.Error("获取连接异常: ", err)
		return err
	}
	return conn.DeletePool(poolName)
}
