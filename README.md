# cloudfoundry-call-kubernetes-example
using Cloud Foundry Service Broker API (v2.6) to expose Kubernetes orchestrated service

## Notification

__How to build and install__

This project is scheduled as a local workspace tree. because all imported package isn't naming github repository, thus 
unable to be build directly.

To resolve this, you may create some symbolic-link on $GOPATH/src, like:

>* `~/workspace$ ln -s $GOPATH/src/github.com/stackdocker/cloudfoundry-call-kubernetes-example/api $GOPATH/src/api`
>* `~/workspace$ ln -s $GOPATH/src/github.com/stackdocker/cloudfoundry-call-kubernetes-example/server $GOPATH/src/server`

从github下载后无法直接build和install，因为import路径没有指向代码在仓库的位置. 
在Linux下可以按上述命令, 为项目每个目录在$GOPATH的src下创建符号链接

__How to verify service wokring__

Using Chrome or other compatible browser to visit _http://<ip addr>:8080/sayhello_. Or in Linux terminal, type:

    $ curl <ip addr>:8080/sayhello
    
and this will simply text out _hello world_