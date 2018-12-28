package main

import (
	"fmt"

	"github.com/fenggolang/client-go-example/common"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientset *kubernetes.Clientset
		podList   *core_v1.PodList
		err       error
	)

	// 初始化k8s客户端
	if clientset, err = common.InitClient(); err != nil {
		fmt.Println(err)
		return
	}

	// 获取default命名空间下的所有pod
	if podList, err = clientset.CoreV1().Pods("wpc").List(meta_v1.ListOptions{}); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*podList)
	return
}
