package main

import (
	"fmt"
	"github.com/ceph/go-ceph/rados"
	"github.com/ceph/go-ceph/rbd"
	"os"
)

const (
	DefaultRadosConfigFile = "/etc/ceph/ceph.conf"
	DefaultBaseImageSize   = 10 * 1024 * 1024 * 1024
	DefaultPoolName        = "rbd"
)

/**
获取连接
 */
func getConnection(monitors, user, key string) (*rados.Conn, error) {
	conn, err := rados.NewConnWithUser(user)
	if err != nil {
		return nil, err
	}
	args := []string{"--client_mount_timeout", "15", "-m", monitors, "--key", key}
	err = conn.ParseCmdLineArgs(args)
	if err != nil {
		return nil, fmt.Errorf("ParseCmdLineArgs: %v", err)
	}
	err = conn.Connect()
	if err != nil {
		return nil, fmt.Errorf("Connect: %v", err)
	}
	return conn, nil
}

/**
文件是否存在
 */
func fileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func main() {
	if fileExist(DefaultRadosConfigFile) {
		//connect to the cluster
		conn, _ := rados.NewConn()
		if err := conn.ReadConfigFile(DefaultRadosConfigFile); err != nil {
			fmt.Println("Rbd read config warn: ", err)
		}
		if err := conn.Connect(); err != nil {
			fmt.Println("Rbd connect failed: ", err)
			return
		}
	}

	conn, err := getConnection("192.168.113.215:6789,192.168.113.216:6789,192.168.113.217:6789", "admin", "AQB+AsFew2rtHRAAwEpQAa1LOG9cYK7k66vtQA==")
	fmt.Println("获取连接成功! GetInstanceID: ", conn.GetInstanceID())
	if err != nil{
		fmt.Println("获取连接异常: ", err)
		return
	}

	pools, err := conn.ListPools()
	fmt.Println("ListPools: ",pools)

	// connect to the pool
	ioctx, err := conn.OpenIOContext(DefaultPoolName)
	if err != nil {
		fmt.Printf("Rbd open pool failed: %v", err)
		return
	}

	// create base image
	baseImageName := "test"
	_, err = rbd.Create(ioctx, baseImageName, DefaultBaseImageSize, 22, rbd.RbdFeatureLayering)
	if err != nil {
		fmt.Printf("Rbd create image failed: %v", err)
		return
	}

	img := rbd.GetImage(ioctx, baseImageName)

	// we should open base image first
	if err := img.Open(); err != nil {
		fmt.Printf("Rbd open image  failed: %v", err)
		return
	}

	defer img.Close()

	// create snapshot
	snapName := "test-snap"
	snapshot, err := img.CreateSnapshot(snapName)
	if err != nil {
		fmt.Printf("Rbd create snapshot failed: %v", err)
		return
	}

	// protect snapshot
	if err := snapshot.Protect(); err != nil {
		fmt.Printf("Rbd create snapshot failed: %v", err)
		return
	}

	// make a clone image based on the snap shot
	cloneImageName := "clone-test"
	_, err = img.Clone(snapName, ioctx, cloneImageName, rbd.RbdFeatureLayering, 22)
	if err != nil {
		fmt.Printf("Rbd clone snapshot failed: %v", err)
		return
	}
	return
}
