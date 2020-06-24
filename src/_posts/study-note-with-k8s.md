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
docker pull kicbase/stable:v0.0.10
```
然后使用minikube指定基础镜像创建集群：
```sh
minikube start --base-image="kicbase/stable:v0.0.10" –image-mirror-country=cn
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

同样的，我将官网的例子替换成了阿里云上的镜像(实际上自己build也行，但是我还不了解部署应用的事情)。
```sh
# create用于创建一系列资源，可以通过yml、json，具体查看k create --help
# deployment是一个二级指令，可以通过指定镜像创建一个自定义名称的deployment，具体查看k create deployment --help
k create deployment hello-minikube --image=registry.cn-hangzhou.aliyuncs.com/google_containers/echoserver:1.4
# deployment.apps/hello-minikube created
k get pod # 查看pod状态
# NAME                              READY   STATUS    RESTARTS   AGE
# hello-minikube-789df8546f-qfrnn   1/1     Running   0          81s
```
值得一提的是，官方文档的例子使用的是`k8s.gcr.io/echoserver:1.4`，而我们可以替换成`registry.cn-hangzhou.aliyuncs.com/google_containers/echoserver:1.4`，换句话说，带`k8s.gcr.io`开头的都可以替换成阿里云镜像地址。

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
## 命名空间演练
命名空间(Namespace)用于逻辑上隔离上下文，方便不同的团队在同一个集群不同作用域中互不干扰的进行作业。

> Kubernetes 命名空间 有助于不同的项目、团队或客户去共享 Kubernetes 集群。
>
> 命名空间命名空间是 Kubernetes 为了在同一物理集群上支持多个虚拟集群而使用的一种抽象。
>
> 名字空间通过以下方式实现这点：
>
> 为名字设置作用域.
> 为集群中的部分资源关联鉴权和策略的机制。

* 创建命名空间
* 为命名空间绑定集群和用户，创建上下文作用域
* 使用上下文切换到对应的命名空间

### 创建命名空间
```yaml
# dev.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: development
  labels:
    name: development
```
编写yaml配置文件，然后使用kubectl读取文件创建，或者也可以直接用命令创建：
```sh
# 读取配置文件
k create -f dev.yaml
# 或者直接用命令创建
k create namespace development 
```
创建完可以通过`k get namespace`查看拥有的命名空间
### 创建上下文并绑定物理集群，指定权限用户
这里使用的是minikube集群，可以通过`k config get clusters`查看。
> ❯ k config get-clusters
> NAME
> 
> minikube

然后使用`set-context`命令绑定该集群，指定用户。
```sh
k config set-context dev --namespace=development \
  --cluster=minikube \
  --user=minikube
```
当然也可以先创建上下文绑定到了命名空间之后使用`set-cluster`来绑定到物理集群。关于用户的后面再提。

使用`k config view`查看物理集群和上下文的信息。
### 切换上下文到对应命名空间
使用`k config current-context`可以查看当前所处上下文，使用`k config use-context dev`将上下文切换到dev中。

在不同的上下文中，创建的资源彼此隔离不可见。当然通信办法是存在的，并且我们可以通过`/proc/$(pid)/ns`来确认lxc命名空间的一些改变。接下来开始学习创建资源的部分，同时可以观察到不同上下文下的隔离性和通信方式。
### 网络策略
minkube使用`flag:--network-plugin`启用对NetworkPolicy的支持[#528](https://github.com/kubernetes/minikube/issues/528)：
```sh
minikube start --network-plugin=cni --enable-default-cni
```
然后安装网络规则驱动，可以选的包括**Calico、Cilium、Kube-router**等。由于笔者使用的k8s版本是1.18.3，前两者还未支持(折腾了两三个小时:sob:)，因此选择[Kube-router](https://github.com/cloudnativelabs/kube-router/blob/master/docs/kubeadm.md)，选择kubeadm安装方式即可：
```sh
k apply -f https://raw.githubusercontent.com/cloudnativelabs/kube-router/master/daemonset/kubeadm-kuberouter.yaml
```
接下来验证效果：
```sh
# 首先创建一个nginx的deployment，暴露成80端口服务
k create deployment nginx --image=nginx
k expose deployment nginx --port=80
# 并使用临时容器busybox访问这个服务
k run busybox --rm -it --image=busybox -- wget --spider --timeout=1 nginx
# Connecting to nginx (10.110.208.165:80)
# remote file exists
```
使用`k describe deployment nginx`来观察镜像nginx创建的容器携带有`app=nginx`的Label。

针对此标签创建一个网络策略yaml文件：
```yaml
# access-nginx.yaml
kind: NetworkPolicy
apiVersion: networking.k8s.io/v1
metadata:
  name: access-nginx
spec:
  podSelector:
    matchLabels:
      app: nginx # 规则对访问app=nginx标签的pod生效
  ingress:
  - from:
    - podSelector:
        matchLabels:
          access: "true" # 规则允许带有access=true标签的pod通过
```
使用`k apply -f access-nginx.yaml`应用规则，再次使用busybox访问发现`wget: download timed out`超时失败。

修改busybox临时容器，使其带有标签`access=true`，再次访问即成功：
```sh
❯ k run busybox --rm -it --labels="access=true" --image=busybox /bin/sh
If you don't see a command prompt, try pressing enter.
/ # wget --spider --timeout=1 nginx
Connecting to nginx (10.110.208.165:80)
remote file exists
```
## 卷
---
先写一点点，未完待续...