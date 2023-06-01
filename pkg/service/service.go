package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	"k8s-baseline-scanner/config"
	p "k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

func (*resource) GetAllService() {
	var (
		svc       p.Service
		asvc      p.AllService
		GetSvcErr = "获取service失败: "
		msg       = "service"
	)

	for _, ns := range tools.GetNamespace() {
		svcList, err := p.K8sInit.GetClientSet().CoreV1().Services(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Error(errors.New(GetSvcErr + err.Error()))
			panic(errors.New(GetSvcErr + err.Error()))
		}
		for _, svc2 := range svcList.Items {
			svc.Selector = make(map[string]string, 10)
			svc.ServiceInfo = ns + "," + svc2.Name + "," + string(svc2.Spec.Type)
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
			asvc.AllService = append(asvc.AllService, svc)
			svc = p.Service{}
		}
	}
	tools.Writer.WriteIntoFile(config.ServiceFile, &asvc)
	filePath := config.DataDirectory + config.ServiceFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
