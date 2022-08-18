# yuurei - 自用个人网站


## 服务器端

在config/application.yml 填写mysql以及redis配置

```
cd server 

go get
go build -o main.exe main.go routes.go
./main.exe
```

## WEB端

```
cd web

npm install 
npm run serve
```
