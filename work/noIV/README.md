# 制作镜像
```
make push
```
# 部署
```
k create -f deploy/deployment.yaml
```
# 从 Promethus 界面中查询延时指标数据
![延时指标数据](image/延时指标数据.png)
# 创建一个 Grafana Dashboard 展现延时分配情况
![延时分配情况](image/延时分配情况.png)