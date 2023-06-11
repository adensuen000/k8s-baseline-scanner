package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	b "k8s-baseline-scanner-v2/base"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PullAllPvc() {
	var (
		pvc     b.Pvc
		apvc    b.AllPvc
		msgName = "pvc"
		file    = b.PvcFile
	)
	apvc = make(map[string]b.Pvc)
	for _, ns := range b.GetNamespaceList() {
		apvcKey := ns
		pvcList, err := b.K8sInit.GetClientSet().CoreV1().PersistentVolumeClaims(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			GetPvcErr := "获取pvc失败: "
			logger.Error(errors.New(GetPvcErr + err.Error()))
			panic(errors.New(GetPvcErr + err.Error()))
		}
		for _, pvc2 := range pvcList.Items {
			info := pvc2.Name + "," + pvc2.Spec.Resources.Requests.Storage().String() + "," + pvc2.Spec.Resources.Limits.Storage().String()
			pvc.PvcInfo = append(pvc.PvcInfo, info)
		}
		apvc[apvcKey] = pvc
		pvc = b.Pvc{}
	}
	b.WriteDataIntoFile(msgName, file, &apvc)
}
