package utils

import (
	"github.com/ceph/go-ceph/rados"
	"github.com/kataras/iris"
	"os"
)

/**
 * 日志模块
 */
var log = iris.New().Logger()

const (
	DefaultRadosConfigFile = "/etc/ceph/ceph.conf"
	DefaultBaseImageSize   = 10 * 1024 * 1024 * 1024
	DefaultPoolName        = "rbd"
)

/**
 * 获取Ceph连接
 */
func GetConnection(monitors, user, key string) (*rados.Conn, error) {
	// 默认通过读取本地配置文件，创建连接
	if FileExist(DefaultRadosConfigFile) {
		//connect to the cluster
		conn, _ := rados.NewConn()
		if err := conn.ReadConfigFile(DefaultRadosConfigFile); err != nil {
			log.Warn("文件/etc/ceph/ceph.conf不存在！", err)
		}
		if err := conn.Connect(); err != nil {
			log.Warn("通过配置文件建立ceph连接失败！", err)
		}
	}
	// 通过用户传递的参数，创建ceph连接
	conn, err := rados.NewConnWithUser(user)
	if err != nil {
		log.Error("建立ceph连接失败！", err)
		return nil, nil
	}
	args := []string{"--client_mount_timeout", "15", "-m", monitors, "--key", key}
	err = conn.ParseCmdLineArgs(args)
	if err != nil {
		log.Error("解析参数失败！", err)
		return nil, nil
	}
	err = conn.Connect()
	if err != nil {
		log.Error("通过传递参数建立ceph连接失败！", err)
		return nil, nil
	}
	return conn, nil
}

/**
 * 文件是否存在
 * return 存在
 */
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return os.IsExist(err)
}
