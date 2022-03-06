## 1. 编译
GOOS=linux GOARCH=amd64 go build -o httpServer main.go
## 2. 上传docker 所在机器

## 3. docker build 
```shell
docker build .  -t 892028617/http-server:202203062200
```

## 4. 推送docker 
```shell
docker push 892028617/http-server:202203062200
```

## 5. 运行docker 
```shell
 docker run -d -p 80:80 892028617/http-server:202203062200
```

## 6. 进入容器
```shell
# 查找票pid
docker inspect -f {{.State.Pid}} 3ed382d5ee65ab4

进入容器
 nsenter --target  93994 --mount --uts --ipc --net --pid
```

