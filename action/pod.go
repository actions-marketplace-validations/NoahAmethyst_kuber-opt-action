package action

import (
	"fmt"
	"kube-opt/client"
	"os"
)

func DelPod(namespace string) {
	var podName string
	if app := os.Getenv("INPUT_APP"); len(app) > 0 {
		println(fmt.Sprintf("Use app:%s", app))
		if resp, err := client.GetPods(namespace); err != nil {
			println(fmt.Sprintf("Get remote kuber pods failed:%s", err.Error()))
			os.Exit(1)
		} else {
			for _, _pod := range resp.Pods {
				if _pod.App == app {
					podName = _pod.PodId
					break
				}
			}
			if len(podName) == 0 {
				println(fmt.Sprintf("No app %s found.", app))
				os.Exit(0)
			}
		}
	} else {
		podName = os.Getenv("INPUT_POD")
		if len(podName) == 0 {
			println("app or pod must specific one of them.")
			os.Exit(1)
		}
		println(fmt.Sprintf("Use pod-name:%s", podName))
	}

	println(fmt.Sprintf("Get pod:%s", podName))

	if resp, err := client.DelPod(namespace, podName); err != nil {
		println(fmt.Sprintf("Delete pod failed:%s", err.Error()))
		os.Exit(1)
	} else if len(resp.Message) > 0 {
		println(fmt.Sprintf("Delete pod failed:%s", resp.Message))
		os.Exit(1)
	}

	println(fmt.Sprintf("Delete pod %s success.", podName))
}
