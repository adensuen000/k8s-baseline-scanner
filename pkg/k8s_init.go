package pkg

import (
	"errors"
	"github.com/wonderivan/logger"
	"k8s-baseline-scanner/pkg/tools"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type k8sInit struct {
}

var K8sInit k8sInit

func (*k8sInit) GetClientSet() *kubernetes.Clientset {
	kubeconfig := tools.GetKubeconfig()
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

func (*k8sInit) GetDynamicClient() *dynamic.DynamicClient {
	kubeconfig := tools.GetKubeconfig()
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		strErr := "获取配置文件失败: "
		logger.Error(errors.New(strErr + err.Error()))
		panic(errors.New(strErr + err.Error()))
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		strErr := "初始化dynamicClient失败: "
		logger.Error(errors.New(strErr + err.Error()))
		panic(errors.New(strErr + err.Error()))
	}
	return dynamicClient
}
