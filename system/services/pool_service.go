package services

import (
	"github.com/ceph/go-ceph/system/datamodels"
	"github.com/ceph/go-ceph/system/utils"
	"github.com/ceph/go-ceph/system/web"
	"github.com/kataras/iris"
)

/**
 * 日志模块
 */
var log = iris.New().Logger()

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
	AddPool(poolName string) error

	/*
		通过poolName删除数据库
	*/
	DeletePoolByName(poolName string) error

	/*
		修改Ceph池配置
	*/
	UpdatePool(pool datamodels.Pool) error
}

// NewPoolService returns the default pool service.
func NewPoolService() IPoolService {
	return &poolService{}
}

type poolService struct {
}

func (p poolService) GetPools(config web.ConnConfig) ([]string, error) {
	log.Info("==>GetPools")

	//获取调go-ceph获取pool
	conn, err := utils.GetConnection(config)
	log.Info("获取连接成功! GetInstanceID: ", conn.GetInstanceID())
	if err != nil {
		log.Error("获取连接异常: ", err)
		return nil, err
	}
	pools, err := conn.ListPools()
	log.Info("Pools: ", pools)
	return pools, err
}

func (p poolService) GetPoolByName(config web.ConnConfig, poolName string) (int64, error) {
	conn, err := utils.GetConnection(config)
	poolId, err := conn.GetPoolByName(poolName)
	return poolId, err
}

func (p poolService) AddPool(poolName string) error {
	panic("implement me")
}

func (p poolService) DeletePoolByName(poolName string) error {
	panic("implement me")
}

func (p poolService) UpdatePool(pool datamodels.Pool) error {
	panic("implement me")
}
