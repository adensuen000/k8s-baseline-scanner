package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	b "k8s-baseline-scanner-v2/base"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

func PullAllService() {
	var (
		svc     b.Service
		svcLt   []b.Service
		asvc    b.AllService
		msgName = "service"
		file    = b.ServiceFile
	)
	asvc = make(map[string][]b.Service)
	for _, ns := range b.GetNamespaceList() {
		svcList, err := b.K8sInit.GetClientSet().CoreV1().Services(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			GetSvcErr := "获取service失败: "
			logger.Error(errors.New(GetSvcErr + err.Error()))
			panic(errors.New(GetSvcErr + err.Error()))
		}
		for _, svc2 := range svcList.Items {
			svc.Selector = make(map[string]string)
			svc.ServiceInfo = svc2.Name + "," + string(svc2.Spec.Type)
			for k, v := range svc2.Spec.Selector {
				svc.Selector[k] = v
			}
			if string(svc2.Spec.Type) == "NodePort" {
				for _, port := range svc2.Spec.Ports {
					portInfo := string(port.Protocol) + "," +
						strconv.Itoa(int(port.NodePort)) + "," +
						strconv.Itoa(int(port.Port)) + "," +
						port.TargetPort.String()
					svc.Potrs = append(svc.Potrs, portInfo)
				}
			} else {
				for _, port := range svc2.Spec.Ports {
					portInfo := string(port.Protocol) + "," +
						strconv.Itoa(int(port.Port)) + "," +
						port.TargetPort.String()
					svc.Potrs = append(svc.Potrs, portInfo)
				}
			}
			svcLt = append(svcLt, svc)
			svc = b.Service{}
		}
		asvcKey := ns
		asvc[asvcKey] = svcLt
		svcLt = nil
	}
	b.WriteDataIntoFile(msgName, file, &asvc)
}
