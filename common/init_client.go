package common

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"io/ioutil"
	"k8s.io/client-go/tools/clientcmd"
)

// 初始化k8s客户端
func InitClient() (clientset *kubernetes.Clientset,err error) {
	var (
		restConf *rest.Config
	)

	if restConf,err =  GetRestConf();err!=nil{
		return
	}
	// 生成clientset配置
	if clientset,err = kubernetes.NewForConfig(restConf);err!=nil{
		return
	}

	return
}

// 获取k8s restful client 配置
func GetRestConf() (restConf *rest.Config, err error) {
	var (
		kubeconfig []byte
	)

	// 读取kubeconfig文件
	if kubeconfig,err = ioutil.ReadFile("./admin.conf");err!=nil{
		return
	}
	// 生成rest client配置
	if restConf,err = clientcmd.RESTConfigFromKubeConfig(kubeconfig);err!=nil{
		return
	}

	return
}