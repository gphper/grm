<p align="center">
  <img src="https://user-images.githubusercontent.com/18718299/176125402-04261517-be75-43a2-8687-3d5e8f9397e9.png" alt="grm" width="140">
</p>

![vue3](http://img.shields.io/badge/vue3-element--plus-blue.svg?style=flat-square) ![go](http://img.shields.io/badge/go-gin-blue.svg?style=flat-square) ![visitor](https://visitor-badge.glitch.me/badge?page_id=gphper.grm)

### 介绍 [Introduction]

基于go+vue的web版redis管理工具【Web redis management tool based on golang and vue】

### 功能清单

- :white_check_mark: 管理连接(直连和SSH)、切换DB

- :white_check_mark: 支持 string/list/set/zset/hash/stream 类型的增删查改

- :white_check_mark: 编译打包成独立的二进制文件

- :white_check_mark: 服务信息展示

- :white_check_mark: 支持命令行

- :white_check_mark: 用户鉴权

- :white_check_mark: 操作日志

- :black_square_button: 支持LUA脚本

## 编译代码
+ web目录下
  `npm run build`
+ grm目录下
  `go build -ldflags "-s -w" .\main.go`
+ 推荐使用upx再次压缩
  `upx -9 main.exe -o grm.exe`  
  
## 运行项目
  * 先添加用户 `grm user add`
  * 执行 `grm run -H ip地址 -p ip端口`  
  * 访问地址 http://ip地址:ip端口//static/#/

## 持续更新中。。。
### 登录
![login](https://user-images.githubusercontent.com/18718299/180608188-9d7a3d97-3c4c-40ea-bcfe-444ed0fc2900.gif)
### 服务管理
![conn](https://user-images.githubusercontent.com/18718299/179389039-d626c654-2874-40e0-951b-27a759d66192.gif)
### 数据展示
![show](https://user-images.githubusercontent.com/18718299/179389052-2229d782-3551-4e07-81e1-6ed8e58d8776.gif)
### Terminal
![cmd](https://user-images.githubusercontent.com/18718299/179389058-039fd95d-3f59-49e2-8141-c994964aa7b0.gif)
### 添加数据
![add](https://user-images.githubusercontent.com/18718299/179389065-98829c30-6098-44de-8471-5ba24be6aab1.gif)

## 支持作者
如果觉得这个项目对你有帮助的话，请留下一颗star ⭐⭐鼓励一下！
If the project is helpful to you, please give a star to encourage me
