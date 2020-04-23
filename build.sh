#打包
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

#打镜像
#docker build -f runing/app.df -t testapp:latest .

#把镜像打tag变成和仓库一致
#docker tag  testapp:lastest 127.0.0.1:5000/myapp:lastest


#推到镜像上去
#docker push 127.0.0.1:5000/myapp:lastest

#起compose
#docker-compose -f runing/docker.yaml pull
#docker-compose -f runing/docker.yaml up
 