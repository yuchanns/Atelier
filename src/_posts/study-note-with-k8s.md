---
title: k8s学习笔记
date: 2020-06-17 23:12:00
category: devops
img: /images/k8s.png
tags:
  - devops
  - k8s
---
k8s学习笔记。
<!-- more -->

![](/images/k8s.png)

先从简单的学起，记录笔记。手册需要好好研究——

## Hello Minikube
### 安装环境
部署配置：
> OS: Deepin 15.11 x86_64
> 
> Model: NUC8i7BEH J72992-307
> 
> Kernel: 4.15.0-30deepin-generic
> 
> CPU: Intel i7-8559U (8) @ 4.5GHz
> 
> GPU: Intel Integrated Graphics
> 
> Memory: 15909MB

通过MBP访问局域网NUC环境。

确保支持虚拟化技术：
```sh
grep -E --color 'vmx|svm' /proc/cpuinfo # 输出不为空
```

### 安装docker
```sh
sudo apt install docker-ce # 通过apt管理包安装docker

sudo groupadd docker # 创建docker组(一般安装docker会自动创建)

sudo gpasswd -a ${USER} docker # 将当前用户加入到docker用户组

sudo service docker restart # 重启docker

docker version
# Docker Engine - Community
# Version: 18.09.6
```

配置`/etc/docker/daemon.json`的镜像源为国内镜像源([阿里地址](https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors))
### 安装kubectl
需要科学上网，我的做法是通过MBP下载二进制包然后再通过FileZilla传递到NUC上。
```sh
# curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl # 可下载最新版
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/linux/amd64/kubectl

chmod +x ./kubectl # 可执行

sudo mv ./kubectl /usr/local/bin/kubectl # 配置环境可用
```
配置shell补全和别名，我使用的是oh-my-zsh+powerlevel10k主题：
```sh
# 在~/.zshrc中
plugins = (kubectl)

alias k=kubectl
complete -F __start_kubectl k

# shell输出
k version
# Client Version: version.Info{Major:"1", Minor:"18", GitVersion:"v1.18.3", GitCommit:"2e7996e3e2712684bc73f0dec0200d64eec7fe40", GitTreeState:"clean", BuildDate:"2020-05-20T12:52:00Z", GoVersion:"go1.13.9", Compiler:"gc", Platform:"linux/amd64"}
# Server Version: version.Info{Major:"1", Minor:"18", GitVersion:"v1.18.3", GitCommit:"2e7996e3e2712684bc73f0dec0200d64eec7fe40", GitTreeState:"clean", BuildDate:"2020-05-20T12:43:34Z", GoVersion:"go1.13.9", Compiler:"gc", Platform:"linux/amd64"}
```
### 安装Minikube
同上，通过二进制文件上传解决：
```sh
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
  && chmod +x minikube

sudo mkdir -p /usr/local/bin/

sudo install minikube /usr/local/bin/ # 配置环境可用
```
### 启动本地单节点k8s集群
需要注意的是，由于网络的障碍，部分资源需要使用国内镜像替代。

首先使用docker拉取基础镜像`kicbase`：
```sh
docker pull kicbase/stable:v0.0.11
```
然后使用minikube指定基础镜像创建集群：
```sh
minikube start --base-image="kicbase/stable:v0.0.11" –image-mirror-country=cn
```
启动时可配置更多参数，使用`minikube start --help`查看。

> 如果你看官方文档或者通过网络途径查看其他人描述的安装心得体会，往往会看到他们说，“设置--image-repository和--image-mirror-country就可以下载”，事实上这两个flag并不是用来拉取base image的，在minikube内部拉取镜像时就有作用了。

启动成功后，可以使用kubectl查看集群状况：
```sh
k cluster-info
# Kubernetes master is running at https://172.17.0.3:8443
# KubeDNS is running at https://172.17.0.3:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

# To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```

访问dashboard图形界面：
```sh
minikube dashboard --url # 添加url是避免弹出浏览器，不加就会弹出，方便本地访问的
# ð¤  正在验证 dashboard 运行情况 ...
# ð  Launching proxy ...
# ð¤  正在验证 proxy 运行状况 ...
# http://127.0.0.1:33081/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/
```
如果在是本机访问，直接访问这个地址就可以了。如果是局域网甚至远程访问，则需要设置转发：
```sh
# address允许访问的ip来源，disable filter取消验证，比较危险，不适合公网使用，更多参数可以通过k proxy --help查看
k proxy --address='0.0.0.0' --disable-filter=true
# W0617 22:17:30.202267   15907 proxy.go:167] Request filter disabled, your proxy is vulnerable to XSRF attacks, please be cautious
# Starting to serve on [::]:8001
```
这时候将域名替换成NUC的局域网ip和对应的端口就可以访问了，比如我的NUC局域网ip是**192.168.199.140**，那么访问`http://192.168.199.140:8001/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/`即可看到dashboard图形化界面。

![](/images/k8s-dashboard.png)

### 部署一个简单的镜像
通过镜像部署一个简单的echoServer，然后使用ssh设置转发，方便局域网访问。

同样的，我将官网的例子替换成了dockerhub上的镜像(实际上自己build也行，但是我还不了解部署应用的事情)。
```sh
# create用于创建一系列资源，可以通过yml、json，具体查看k create --help
# deployment是一个二级指令，可以通过指定镜像创建一个自定义名称的deployment，具体查看k create deployment --help
k create deployment hello-minikube --image=jmalloc/echo-server
# deployment.apps/hello-minikube created
k get pod # 查看pod状态
# NAME                              READY   STATUS    RESTARTS   AGE
# hello-minikube-789df8546f-qfrnn   1/1     Running   0          81s
```
部署完成功后，为了能够访问，需要暴露为Service：
```sh
# 指定Service类型为NodePort，并且指定服务端口为8080
k expose deployment hello-minikube --type=NodePort --port=8080
# service/hello-minikube exposed
```
然后获取Service的地址：
```sh
minikube service hello-minikube --url
# http://172.17.0.3:31869
```
同样的，仅限于本地访问，如果要通过局域网访问，需要使用ssh转发：
```sh
# 这里指定任意来源ip访问9000端口就转发到Service的地址
# 详细说明可以通过man ssh来查看
# 摘抄：
# -L [bind_address:]port:remote_socket
# 后面那个地址则是ssh要登录的服务器ip地址
ssh -L 0.0.0.0:9000:172.17.0.3:31869 192.168.199.140
```
这时候我们通过浏览器访问NUC的ip:9000就可以访问到这个服务了。

![](/images/k8s-echoServer.png)

### 删除清理
```sh
# 查看正在运行的Services
k get service
# 删除暴露的Service
k delete svc hello-minikube
# 删除deployment
k delete deployment hello-minikube
# 取消转发
netstat -tunlp | grep 9000
# tcp        0      0 0.0.0.0:9000            0.0.0.0:*               LISTEN      8566/ssh
kill -2 8566 # 删掉转发的进程
# 使用Ctrl+C退出k proxy转发和k dashboard进程
# 停止集群
minikube stop
# 删除集群
minikube delete
```
---
先写一点点，未完待续...