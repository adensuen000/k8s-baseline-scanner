package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	"k8s-baseline-scanner/config"
	p "k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (*resource) GetAllPvc() {
	var (
		pvc       p.Pvc
		apvc      p.Allpvc
		GetPvcErr = "获取pvc失败: "
		msg       = "pvc"
	)

	for _, ns := range tools.GetNamespace() {
		pvcList, err := p.K8sInit.GetClientSet().CoreV1().PersistentVolumeClaims(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Error(errors.New(GetPvcErr + err.Error()))
			panic(errors.New(GetPvcErr + err.Error()))
		}
		for _, pvc2 := range pvcList.Items {
			info := pvc2.Namespace + "," + pvc2.Name + "," + pvc2.Spec.Resources.Requests.Storage().String() + "," + pvc2.Spec.Resources.Limits.Storage().String()
			pvc.PvcInfo = append(pvc.PvcInfo, info)
		}
		apvc.AllPvc = append(apvc.AllPvc, pvc)
		pvc = p.Pvc{}
	}
	tools.Writer.WriteIntoFile(config.PvcFile, &apvc)
	filePath := config.DataDirectory + config.PvcFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
