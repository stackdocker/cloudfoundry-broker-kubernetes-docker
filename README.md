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

__Run in debug mode__

Like:

>`$ go run main.go --redis-addr 192.168.0.1:6379 --sentinel-addrs 192.168.0.11:26379,192.168.0.12:26379`

__Test with curl__

Like [request to binding](http://docs.cloudfoundry.org/services/api-v2.6.html#binding):

>`$ curl -i http://127.0.0.1:8080/v2/service_instances/00000000-0000-0000-0000-000000000000/service_bindings/00000000-0000-0000-0000-000000000001 -d '{"service_id": "9c372bbc-1e7b-472b-bcb6-eeda5b21eb35", "plan_id": "8cfbbaf5-efdb-41c1-89ab-f797185f7818", "app_guid": "00000000-0000-0000-0000-000000000002", "parameters": {"organization_guid": "a924baea-363a-457e-8708-0eda152e76c5", "space_guid": "1fde0510-cc7e-42c0-8321-1505d51f0c29", "parameter1": 1, "parameter2": "value"}}' -X PUT -H "X-Broker-API-Version: 2.6" -H "Content-Type: application/json"; echo`

>`HTTP/1.1 200 OK`

>`Content-Type: application/json`

>`Date: Tue, 20 Oct 2015 14:01:12 GMT`

>`Content-Length: 83`

>`{"credentials":{"Address":"192.168.0.1:6379","Db":0,"Network":"tcp","Password":""}}`


__Verify service wokring__

Using Chrome or other compatible browser to visit _http://<ip addr>:8080/sayhello_. Or in Linux terminal, type:

    $ curl <ip addr>:8080/sayhello
    
and this will simply text out _hello world_