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

- :black_square_button: 支持命令行

- :black_square_button: 支持LUA脚本

- :black_square_button: 操作日志

- :black_square_button: 用户鉴权


## 编译代码
+ web目录下
  `npm run build`
+ grm目录下
  `go build -ldflags "-s -w" .\main.go`
+ 推荐使用upx再次压缩
  `upx -9 main.exe -o grm.exe`  

## 持续更新中。。。

![show1](https://user-images.githubusercontent.com/18718299/177555667-9e60ab58-6483-4d33-915a-b89c1c262a93.gif)

![show](https://user-images.githubusercontent.com/18718299/176183368-44597b01-977b-44c4-bd1b-b7c987f1e6c4.gif)

## 支持作者
如果觉得这个项目对你有帮助的话，请留下一颗star ⭐⭐鼓励一下！
If the project is helpful to you, please give a star to encourage me
