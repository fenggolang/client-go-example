package main

import (
	"k8s.io/client-go/kubernetes"
	"github.com/fenggolang/client-go-example/common"
	"fmt"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	core_v1 "k8s.io/api/core/v1"
)

func main() {
	var (
		clientset *kubernetes.Clientset
		podList *core_v1.PodList
		err error
	)

	// 初始化k8s客户端
	if clientset,err = common.InitClient();err!=nil{
		fmt.Println(err)
		return
	}

	// 获取default命名空间下的所有pod
	if podList,err = clientset.CoreV1().Pods("default").List(meta_v1.ListOptions{});err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(*podList)
	return
}