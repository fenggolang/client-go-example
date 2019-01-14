package main

import (
	"fmt"

	"github.com/fenggolang/client-go-example/common"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	var (
		clientset     *kubernetes.Clientset
		tailLines     int64
		req           *rest.Request
		result        rest.Result
		logs          []byte
		err           error
		podName       = "nfs-client-provisioner-69665d796c-sxgwq"
		containerName = "nfs-client-provisioner"
	)

	// 初始化k8s客户端
	if clientset, err = common.InitClient(); err != nil {
		fmt.Println(err)
		return
	}

	// 生成获取POD日志请求
	req = clientset.CoreV1().Pods("default").GetLogs(podName, &v1.PodLogOptions{Container: containerName, TailLines: &tailLines})

	//req.Stream() // 也可以实现Do的效果

	// 发送请求
	if result = req.Do(); result.Error() != nil {
		err = result.Error()
		fmt.Println(err)
		return
	}

	// 获取结果
	if logs, err = result.Raw(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("容器输出：", string(logs))
}
