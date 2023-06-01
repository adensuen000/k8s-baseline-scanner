package pkg

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	"k8s-baseline-scanner/pkg/tools"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var K8sRce k8sRce

type k8sRce struct {
}

func (*k8sRce) GetStaticPod() {

}

// 获取deployment、statefulset、daemonset资源对应的容器
func (*k8sRce) GetContainers() AllContainers {
	var (
		resourcectr       ResourceContainer
		allContainers     AllContainers
		getDeploymentsErr = "获取deployments失败: "
		getStsErr         = "获取sts失败: "
		getDsErr          = "获取ds失败: "
		resourceType      string
		namespace         string
	)
	for _, ns := range tools.GetNamespace() {
		//获取deployment类型的应用和对应的container
		deployList, err := K8sInit.GetClientSet().AppsV1().Deployments(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Error(errors.New(getDeploymentsErr + err.Error()))
			panic(errors.New(getDeploymentsErr + err.Error()))
		}
		for _, deploy := range deployList.Items {
			namespace = ns
			resourceType = "Deployment"
			resourcectr.ResourceInfo = namespace + "," + resourceType + "," + deploy.Name
			resourcectr.Containers = deploy.Spec.Template.Spec.Containers
			allContainers.AllContainers = append(allContainers.AllContainers, resourcectr)
		}

		//获取sts类型的应用和对应的container
		stsList, err := K8sInit.GetClientSet().AppsV1().StatefulSets(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Error(errors.New(getStsErr + err.Error()))
			panic(errors.New(getStsErr + err.Error()))
		}
		for _, sts := range stsList.Items {
			namespace = ns
			resourceType = "StatefulSet"
			resourcectr.ResourceInfo = namespace + "," + resourceType + "," + sts.Name
			resourcectr.Containers = sts.Spec.Template.Spec.Containers
			allContainers.AllContainers = append(allContainers.AllContainers, resourcectr)
		}

		//获取ds类型的应用和对应的container
		dsList, err := K8sInit.GetClientSet().AppsV1().StatefulSets(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Error(errors.New(getDsErr + err.Error()))
			panic(errors.New(getDsErr + err.Error()))
		}
		for _, ds := range dsList.Items {
			namespace = ns
			resourceType = "DaemontSet"
			resourcectr.ResourceInfo = namespace + "," + resourceType + "," + ds.Name
			resourcectr.Containers = ds.Spec.Template.Spec.Containers
			allContainers.AllContainers = append(allContainers.AllContainers, resourcectr)
		}
	}
	return allContainers
}

// 获取deployment、statefulset、daemonset资源

func GetResources() []DeployStsDsBase {
	var (
		getDeploymentsErr = "获取deployments失败: "
		getStsErr         = "获取sts失败: "
		getDsErr          = "获取ds失败: "
		dsdb              DeployStsDsBase
		dsdbSce           []DeployStsDsBase
		resourceType      string
		namespace         string
		Resourcename      string
	)

	for _, ns := range tools.GetNamespace() {
		namespace = ns
		//获取deployment类型的应用
		deployList, err := K8sInit.GetClientSet().AppsV1().Deployments(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Error(errors.New(getDeploymentsErr + err.Error()))
			panic(errors.New(getDeploymentsErr + err.Error()))
		}
		for _, deploy := range deployList.Items {
			resourceType = "Deployment"
			Resourcename = deploy.Name
			dsdb.ResourceInfo = namespace + "," + resourceType + "," + Resourcename
			dsdb.Replicas = *deploy.Spec.Replicas
			dsdb.RestartPolicy = string(deploy.Spec.Template.Spec.RestartPolicy)
			dsdb.DnsPolicy = string(deploy.Spec.Template.Spec.DNSPolicy)
			dsdb.Volumes = deploy.Spec.Template.Spec.Volumes
			dsdb.Affinity = deploy.Spec.Template.Spec.Affinity
			dsdb.NodeSelector = deploy.Spec.Template.Spec.NodeSelector
			dsdb.ServiceAccountName = deploy.Spec.Template.Spec.ServiceAccountName
			dsdb.Tolerations = deploy.Spec.Template.Spec.Tolerations
			dsdbSce = append(dsdbSce, dsdb)
			dsdb = DeployStsDsBase{}
		}

		//获取sts类型的应用
		stsList, err := K8sInit.GetClientSet().AppsV1().StatefulSets(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Error(errors.New(getStsErr + err.Error()))
			panic(errors.New(getStsErr + err.Error()))
		}
		for _, sts := range stsList.Items {
			resourceType = "StatefulSet"
			Resourcename = sts.Name
			dsdb.ResourceInfo = namespace + "," + resourceType + "," + Resourcename
			dsdb.Replicas = *sts.Spec.Replicas
			dsdb.RestartPolicy = string(sts.Spec.Template.Spec.RestartPolicy)
			dsdb.DnsPolicy = string(sts.Spec.Template.Spec.DNSPolicy)
			dsdb.Volumes = sts.Spec.Template.Spec.Volumes
			dsdb.Affinity = sts.Spec.Template.Spec.Affinity
			dsdb.NodeSelector = sts.Spec.Template.Spec.NodeSelector
			dsdb.ServiceAccountName = sts.Spec.Template.Spec.ServiceAccountName
			dsdb.Tolerations = sts.Spec.Template.Spec.Tolerations
			dsdbSce = append(dsdbSce, dsdb)
			dsdb = DeployStsDsBase{}
		}

		//获取ds类型的应用
		dsList, err := K8sInit.GetClientSet().AppsV1().StatefulSets(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Error(errors.New(getDsErr + err.Error()))
			panic(errors.New(getDsErr + err.Error()))
		}
		for _, ds := range dsList.Items {
			resourceType = "DaemonSet"
			Resourcename = ds.Name
			dsdb.ResourceInfo = namespace + "," + resourceType + "," + Resourcename
			dsdb.Replicas = *ds.Spec.Replicas
			dsdb.RestartPolicy = string(ds.Spec.Template.Spec.RestartPolicy)
			dsdb.DnsPolicy = string(ds.Spec.Template.Spec.DNSPolicy)
			dsdb.Volumes = ds.Spec.Template.Spec.Volumes
			dsdb.Affinity = ds.Spec.Template.Spec.Affinity
			dsdb.NodeSelector = ds.Spec.Template.Spec.NodeSelector
			dsdb.ServiceAccountName = ds.Spec.Template.Spec.ServiceAccountName
			dsdb.Tolerations = ds.Spec.Template.Spec.Tolerations
			dsdbSce = append(dsdbSce, dsdb)
			dsdb = DeployStsDsBase{}
		}
	}
	return dsdbSce
}
