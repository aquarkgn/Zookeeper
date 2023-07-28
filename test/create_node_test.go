package test

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func TestCreateNode(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		createNode(i, &wg)
		time.Sleep(time.Second * 3)
	}

	wg.Wait()
	fmt.Println("创建完成")
}

func createNode(nodeID int, wg *sync.WaitGroup) {
	defer wg.Done()
	servers := []string{"localhost:2181"} // Zookeeper 服务器地址根据实际情况修改
	nodePrefix := "/node"                 // 节点前缀

	conn, _, err := zk.Connect(servers, time.Second*5)
	if err != nil {
		fmt.Println("连接失败: ", err.Error())
		return
	}
	log.Printf("连接成功: %v\n", servers)
	defer conn.Close()

	path := fmt.Sprintf("%s/%d", nodePrefix, nodeID)
	data := []byte(fmt.Sprintf("Data for node %d", nodeID))

	log.Printf("创建节点: %s\n", path)
	_, err = conn.Create(path, data, 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		fmt.Printf("创建节点失败: %s, %s\n", path, err.Error())
	} else {
		log.Printf("创建节点成功: %s\n", path)
	}
}
