# 一键打包

通过 shell 脚本，下载 Github 仓库代理，本地编译后，打包为 Docker 镜像

# 第一步：添加 Github SSH key

- 生成公钥

`ssh-keygen -t rsa -C '921229265@qq.com'`

- 复制公钥

`cat ~/.ssh/id_rsa.pub`

## 打开 Github 添加 SSH key

https://github.com/settings/ssh/new

# 第二步：配置 go 环境

https://go.dev/doc/install

- 下载 go 执行文件

`wget https://dl.google.com/go/go1.20.4.linux-amd64.tar.gz`

- 解压 go 执行文件

`rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.4.linux-amd64.tar.gz`

- 添加执行文件到环境变量

`echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.profile`

`source $HOME/.profile`

- 更换包管理地址

`go env -w GOPROXY=https://goproxy.cn`

# 第三步：执行脚本

`chmod a+x start.sh`

`./start.sh`

# 脚本内容如下

```shell
#!/bin/bash

# 清理旧文件
echo "********** 第 1 步：准备工作 **********"
rm -rf ~/ipnas6/build
mkdir -p ~/ipnas6/build
cd ~/ipnas6/build
pwd

echo "********** 第 2 步：克隆代码 **********"
git clone git@github.com:gokingliu/ipnas6.git

echo "********** 第 3 步：打包服务端 **********"
cd ipnas6
git pull
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build
cp ipnas6 app/

echo "********** 第 4 步：清理旧镜像 **********"
docker image rm ipnas6:1

echo "********** 第 5 步：打包新镜像 **********"
chmod a+x app/*
docker build --platform=linux/arm64 -t ipnas6:1 .

echo "********** 第 6 步：推送镜像 **********"
docker tag ipnas6:1 crotaliu/ipnas6:arm64
docker push crotaliu/ipnas6:arm64

echo "********** 结束 **********"
```
