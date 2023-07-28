# Zookeeper
此项目是Zookeeper的数据写入测试项目，目的是为了给Grafana提供数据指标。
This project is a data write test project for Zookeeper, aiming to provide data metrics for Grafana.

## 启动ActiveMQ
```shell
cd deployment/docker-compose
docker-compose up -d

cd conf && echo "4lw.commands.whitelist=*" >> zoo.cfg

docker-compose restart
```

