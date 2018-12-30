package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/fenggolang/client-go-example/common"
	core_v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientset     *kubernetes.Clientset
		deployYaml    []byte
		deployJson    []byte
		deployment    = v1beta1.Deployment{}
		k8sDeployment *v1beta1.Deployment
		podList       *core_v1.PodList
		pod           core_v1.Pod
		err           error
	)

	// 初始化k8s客户端
	if clientset, err = common.InitClient(); err != nil {
		fmt.Println("初始化k8s客户端失败:", err)
		return
	}

	// 读取YAML
	if deployYaml, err = ioutil.ReadFile("./nginx.yaml"); err != nil {
		fmt.Println("读取yaml失败：", err)
		return
	}

	// YAML转JSON
	if deployJson, err = yaml.ToJSON(deployYaml); err != nil {
		fmt.Println("yaml转json失败：", err)
		return
	}

	// 反序列化：JSON转struct
	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		fmt.Println("反序列化失败:", err)
		return
	}

	// 给Pod添加label
	deployment.Spec.Template.Labels["deploy_time"] = strconv.Itoa(int(time.Now().Unix()))

	// 更新deployment
	if _, err = clientset.ExtensionsV1beta1().Deployments("wpc").Update(&deployment); err != nil {
		fmt.Println("更新deployment失败：", err)
		return
	}

	// 等待更新完成
	for {
		// 获取k8s中deployment的状态
		if k8sDeployment, err = clientset.ExtensionsV1beta1().Deployments("wpc").Get(deployment.Name, v1.GetOptions{}); err != nil {
			goto RETRY // 休眠1秒再重试
		}

		// 进行状态判定
		if k8sDeployment.Status.UpdatedReplicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.Replicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.AvailableReplicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.ObservedGeneration == k8sDeployment.Generation {
			// 滚动升级完成
			break
		}

		// 打印工作中的pod比例
		fmt.Printf("部署中：(%d/%d)\n", k8sDeployment.Status.AvailableReplicas, *(k8sDeployment.Spec.Replicas))

	RETRY:
		time.Sleep(1 * time.Second)
	}

	fmt.Println("部署成功!")

	// 打印每个pod的状态(可能会打印出terminating中的pod,但最终只会展示新pod列表)
	if podList, err = clientset.CoreV1().Pods("wpc").List(v1.ListOptions{LabelSelector: "app=nginx", IncludeUninitialized: false}); err == nil {
		for _, pod = range podList.Items {
			podName := pod.Name
			podStatus := string(pod.Status.Phase)

			// PodRunning表示pod已绑定到节点，并且所有容器都已启动。
			// 至少有一个容器仍在运行或正在重新启动。
			if podStatus == string(core_v1.PodRunning) {
				// 汇总错误原因不为空
				if pod.Status.Reason != "" {
					podStatus = pod.Status.Reason
					goto KO
				}

				// condition有错误信息
				for _, cond := range pod.Status.Conditions {
					if cond.Type == core_v1.PodReady { // pod就绪状态
						if cond.Status != core_v1.ConditionTrue { // 失败
							podStatus = cond.Reason
						}
						goto KO
					}
				}

				// 没有ready condition, 状态未知
				podStatus = "Unknown"
			}

		KO:
			fmt.Printf("[name:%s status:%s]\n", podName, podStatus)

		}
	}
}
