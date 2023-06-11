package base

import (
	"errors"
	"github.com/wonderivan/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type k8sInit struct {
}

var K8sInit k8sInit

func (*k8sInit) GetClientSet() *kubernetes.Clientset {
	kubeconfig := GetKubeconfig()
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		strErr := "获取配置文件失败: "
		logger.Error(errors.New(strErr + err.Error()))
		panic(errors.New(strErr + err.Error()))
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		strErr := "初始化clientSet失败: "
		logger.Error(errors.New(strErr + err.Error()))
		panic(errors.New(strErr + err.Error()))
	}
	return clientSet
}
