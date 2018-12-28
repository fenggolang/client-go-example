package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fenggolang/client-go-example/common"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientset  *kubernetes.Clientset
		deployYaml []byte
		deployJson []byte
		deployment = v1beta1.Deployment{}
		replicas   int32
		err        error
	)

	// 初始化k8s客户端
	if clientset, err = common.InitClient(); err != nil {
		fmt.Println(err)
		return
	}

	// 读取YAML
	if deployYaml, err = ioutil.ReadFile("./nginx.yaml"); err != nil {
		fmt.Println(err)
		return
	}

	// YAML转JSON
	if deployJson, err = yaml.ToJSON(deployYaml); err != nil {
		fmt.Println(err)
		return
	}

	// JSON转struct
	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		fmt.Println(err)
		return
	}

	// 修改replicas数量为1
	replicas = 1
	deployment.Spec.Replicas = &replicas

	// 查询k8s是否有该deployment
	if _, err = clientset.ExtensionsV1beta1().Deployments("wpc").Get(deployment.Name, meta_v1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			fmt.Println(err)
			return
		}
		// 不存在则创建
		if _, err = clientset.ExtensionsV1beta1().Deployments("wpc").Create(&deployment); err != nil {
			fmt.Println(err)
			return
		}
	} else { // 已存在则更新
		if _, err = clientset.ExtensionsV1beta1().Deployments("wpc").Update(&deployment); err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("apply成功!")
	return
}
