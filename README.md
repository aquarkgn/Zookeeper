# Zookeeper
此项目是Zookeeper的数据写入测试项目，目的是为了给Grafana提供数据指标。
This project is a data write test project for Zookeeper, aiming to provide data metrics for Grafana.

## 启动Zookeeper
```shell
cd deployment/docker-compose
docker-compose up -d

cd conf && echo "4lw.commands.whitelist=*" >> zoo.cfg

docker-compose restart
```

## 启动Zookeeper
```shell
cd test

#创建zookeeper的节点
nohub go test -run TestCreateNode -v > create.log 2>&1 &

#写入zookeeper的节点数据
nohub go test -run TestWriteDataInNode -v > write.log 2>&1 &

#创建zookeeper的节点观察值
nohub go test -run TestObserver -v > watcher.log 2>&1 &
```
