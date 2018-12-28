package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fenggolang/client-go-example/common"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientset      *kubernetes.Clientset
		deployYaml     []byte
		deployJson     []byte
		deployment     = v1beta1.Deployment{}
		containers     []v1.Container
		nginxContainer v1.Container
		err            error
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

	// 定义的container
	nginxContainer.Name = "nginx"
	nginxContainer.Image = "nginx:1.13.8"
	nginxContainer.Resources.Limits.Cpu().Set(1)
	nginxContainer.Resources.Limits.Memory().Set(1)
	nginxContainer.Resources.Requests.Cpu().Set(1)
	nginxContainer.Resources.Requests.Memory().Set(1)
	//nginxContainer.Resources.Limits.Cpu().String()="200m"
	//nginxContainer.Resources.Limits.Memory().String() = "256Mi"
	//nginxContainer.Resources.Requests.Cpu().String() = "100m"
	//nginxContainer.Resources.Requests.Memory().String() = "256Mi"
	// 这样直接追加的话会把整个containers下的字段给替换掉
	containers = append(containers, nginxContainer)

	// 修改podTemplate,定义container列表
	deployment.Spec.Template.Spec.Containers = containers
	//clientset.ExtensionsV1beta1().Deployments("wpc").Patch(&deployment)

	// 更新deployment
	if _, err = clientset.ExtensionsV1beta1().Deployments("wpc").Update(&deployment); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("apply成功!")
	return
}
