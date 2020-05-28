package services

import (
	"github.com/ceph/go-ceph/system/pool/datamodels"
	"github.com/ceph/go-ceph/system/utils"
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
	GetPools() ([]string, error)

	/*
		/api/pool/{pool_name}
	*/
	GetPoolByName(poolName string) ([]datamodels.Pool, error)

	/*
		创建ceph存储池
		/api/pool
	*/
	AddPool(pool datamodels.Pool) (bool, error)

	/*
		通过poolName删除数据库
	*/
	DeletePoolByName(poolName string) (bool, error)

	/*
		修改Ceph池配置
	*/
	UpdatePool(pool datamodels.Pool) (datamodels.Pool, error)
}

// NewPoolService returns the default pool service.
func NewPoolService() IPoolService {
	return &poolService{}
}

type poolService struct {
}

func (p poolService) GetPools() ([]string, error) {
	log.Info("==>GetPools")

	//获取调go-ceph获取pool
	conn, err := utils.GetConnection("192.168.113.215:6789,192.168.113.216:6789,192.168.113.217:6789", "admin", "AQB+AsFew2rtHRAAwEpQAa1LOG9cYK7k66vtQA==")
	log.Info("获取连接成功! GetInstanceID: ", conn.GetInstanceID())
	if err != nil {
		log.Error("获取连接异常: ", err)
		return nil, err
	}
	pools, err := conn.ListPools()
	log.Info("Pools: ", pools)
	return pools, err
}

func (p poolService) GetPoolByName(poolName string) ([]datamodels.Pool, error) {
	panic("implement me")
}

func (p poolService) AddPool(pool datamodels.Pool) (bool, error) {
	panic("implement me")
}

func (p poolService) DeletePoolByName(poolName string) (bool, error) {
	panic("implement me")
}

func (p poolService) UpdatePool(pool datamodels.Pool) (datamodels.Pool, error) {
	panic("implement me")
}
