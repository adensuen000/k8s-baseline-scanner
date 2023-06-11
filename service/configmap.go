package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	b "k8s-baseline-scanner-v2/base"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func PullAllConfigmap() {
	var (
		cm      b.Configmap
		rcm     b.ResourceConfigmap
		acm     b.AllConfigmap
		msgName = "configmap"
		file    = b.ConfigmapFile
	)

	acm = make(map[string]b.ResourceConfigmap)
	for _, rce := range b.GetResources() {
		acmKey := rce.ResourceInfo
		namespace := strings.Split(rce.ResourceInfo, ",")[0]
		if rce.Volumes != nil {
			for _, v := range rce.Volumes {
				if v.ConfigMap != nil {
					cm.CmData = make(map[string]string, 10)
					cm.CmName = v.ConfigMap.Name
					cm2, err := b.K8sInit.GetClientSet().CoreV1().ConfigMaps(namespace).Get(context.TODO(), cm.CmName, metav1.GetOptions{})
					if err != nil {
						GetCmErr := "获取configmap失败: "
						logger.Error(errors.New(GetCmErr + err.Error()))
						panic(errors.New(GetCmErr + err.Error()))
					}
					for k, ve := range cm2.Data {
						data := strings.ReplaceAll(ve, "\n", "\r\n")
						cm.CmData[k] = data
						//cm.CmData[k] = ve
					}
					rcm.ResourceConfigmap = append(rcm.ResourceConfigmap, cm)
					cm = b.Configmap{}
				}
			}
		}
		acm[acmKey] = rcm
		rcm = b.ResourceConfigmap{}
	}
	b.WriteDataIntoFile(msgName, file, &acm)
}
