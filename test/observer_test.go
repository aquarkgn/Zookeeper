package test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func TestObserver(t *testing.T) {
	// 设置 ZooKeeper 连接信息
	zkHosts := []string{"localhost:2181"}
	zkPath := "/zookeeper"

	// 创建多个观察者连接
	numObservers := 5
	for i := 1; i <= numObservers; i++ {
		// 创建一个新的 ZooKeeper 连接
		conn, _, err := zk.Connect(zkHosts, time.Second)
		if err != nil {
			log.Fatalf("Failed to connect to ZooKeeper server: %v", err)
		}

		// 创建一个观察者
		observer := &myWatcher{}

		// 获取节点的值并设置观察者
		data, _, watch, err := conn.GetW(zkPath)
		if err != nil {
			log.Fatalf("Failed to get data from node: %v", err)
		}
		fmt.Printf("Observer %d - Initial data: %s\n", i, string(data))
		go watchNode(observer, watch)
	}

	// 暂停程序，保持观察者连接
	select {}
}

type myWatcher struct{}

// 监听节点变化的回调函数
func (w *myWatcher) Process(event zk.Event) {
	fmt.Printf("Node Event - Type: %s, Path: %s\n", event.Type, event.Path)
}

// 监听节点的变化
func watchNode(observer *myWatcher, watch <-chan zk.Event) {
	// 处理节点变化事件
	for {
		select {
		case event := <-watch:
			observer.Process(event)
		}
	}
}
