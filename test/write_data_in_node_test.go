package test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func TestWriteDataInNode(t *testing.T) {
	// 设置 ZooKeeper 连接信息
	zkHosts := []string{"localhost:2181"}
	zkPath := "/zookeeper/1"

	// 创建 ZooKeeper 连接
	conn, _, err := zk.Connect(zkHosts, time.Second)
	if err != nil {
		log.Fatalf("Failed to connect to ZooKeeper server: %v", err)
	}

	// 持续不断地写入数据
	i := 0
	for {
		data := fmt.Sprintf("写入数字: %v", i)
		err := writeToNode(conn, zkPath, data)
		if err != nil {
			log.Printf("node %s 写入数据失败: %s", zkPath, err.Error())
			break
		} else {
			log.Printf("node %s 写入数据: %s", zkPath, data)
		}

		i++
		// 等待 1 秒后继续写入
		time.Sleep(time.Second)
	}
}

// 将数据写入节点
func writeToNode(conn *zk.Conn, path string, data string) error {
	// 创建或更新节点
	_, err := conn.Set(path, []byte(data), -1)
	if err != nil {
		return err
	}
	return nil
}
