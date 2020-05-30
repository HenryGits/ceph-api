package main

import (
	"fmt"
	"github.com/ceph/go-ceph/rados"
	"github.com/ceph/go-ceph/rbd"
	"github.com/ceph/go-ceph/system/utils"
	"github.com/ceph/go-ceph/system/web"
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

	conf := web.ConnConfig{Key: "AQB+AsFew2rtHRAAwEpQAa1LOG9cYK7k66vtQA==",
		Monitors: "192.168.113.215:6789,192.168.113.216:6789,192.168.113.217:6789", User: "admin"}
	conn, err := utils.GetConnection(conf)
	fmt.Println("获取连接成功! GetInstanceID: ", conn.GetInstanceID())
	if err != nil {
		fmt.Println("获取连接异常: ", err)
		return
	}

	//
	//================================================= pool池操作= =========================================================
	//
	// 获取
	poolId, err := conn.GetPoolByName(DefaultPoolName)
	fmt.Println("poolName==> ", poolId)
	poolStr, _ := conn.GetPoolByID(poolId)
	fmt.Println("poolStr==> ", poolStr)
	// 创建
	conn.MakePool("test-pool")
	//池列表
	pools, err := conn.ListPools()
	fmt.Println("ListPools: ", pools)
	// 删除
	conn.DeletePool("test-pool")

	//
	//================================================= rbd操作= =========================================================
	//
	// connect to the pool
	ioctx, err := conn.OpenIOContext(DefaultPoolName)
	if err != nil {
		fmt.Printf("Rbd open pool failed: %v", err)
		return
	}
	fmt.Println("ioctx conn success...")

	baseImageName := "test2"
	cloneImageName := "clone-test"

	_ = rbd.RemoveImage(ioctx, cloneImageName)
	img, _ := rbd.OpenImage(ioctx, baseImageName, "")
	//// 需要删除快照
	//snap:=img.GetSnapshot(cloneImageName)
	//snap.Remove()
	// 再删除image
	_ = rbd.TrashRemove(ioctx, baseImageName, true)
	err = rbd.RemoveImage(ioctx, baseImageName)
	fmt.Println("RemoveImage==>", err)

	// img 列表
	imgs, _ := rbd.GetImageNames(ioctx)
	fmt.Println("imgs==>", imgs)

	// create base image
	_, err = rbd.Create2(ioctx, baseImageName, DefaultBaseImageSize, rbd.FeatureLayering, 22)
	if err != nil {
		fmt.Printf("Rbd create image failed: %v", err)
		return
	}

	img = rbd.GetImage(ioctx, baseImageName)
	imgSize, _ := img.GetSize()
	fmt.Println("img==>", img)
	fmt.Println("imgSize==>", imgSize)

	// we should open base image first
	if err := img.Open(); err != nil {
		fmt.Printf("Rbd open image  failed: %v", err)
		return
	}

	//// create snapshot
	//snapName := "test-snap"
	//snapshot, err := img.CreateSnapshot(snapName)
	//if err != nil {
	//	fmt.Printf("Rbd create snapshot failed: %v", err)
	//	return
	//}
	//
	//// 锁定 snapshot
	//if err := snapshot.Protect(); err != nil {
	//	fmt.Printf("Rbd create snapshot failed: %v", err)
	//	return
	//}
	//
	//// make a clone image based on the snap shot
	//_, err = img.Clone(snapName, ioctx, cloneImageName, rbd.RbdFeatureLayering, 22)
	//if err != nil {
	//	fmt.Printf("Rbd clone snapshot failed: %v", err)
	//	return
	//}
	//
	//img.Remove()
	//
	//imgs,_ = rbd.GetImageNames(ioctx)
	//fmt.Println("imgs==>", imgs)
}
