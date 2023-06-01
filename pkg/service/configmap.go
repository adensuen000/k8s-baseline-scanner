package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	"k8s-baseline-scanner/config"
	p "k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func (*resource) GetAllConfigmap() {
	var (
		cm       p.Configmap
		rcm      p.ResourceConfigmap
		acm      p.AllConfigmap
		GetCmErr = "获取configmap失败: "
		msg      = "configmap"
	)

	for _, rce := range p.GetResources() {
		rcm.ResourceInfo = rce.ResourceInfo
		namespace := strings.Split(rce.ResourceInfo, ",")[0]
		if rce.Volumes != nil {
			for _, v := range rce.Volumes {
				if v.ConfigMap != nil {
					cm.CmData = make(map[string]string, 10)
					cm.CmName = v.ConfigMap.Name
					cm2, err := p.K8sInit.GetClientSet().CoreV1().ConfigMaps(namespace).Get(context.TODO(), cm.CmName, metav1.GetOptions{})
					if err != nil {
						logger.Error(errors.New(GetCmErr + err.Error()))
						panic(errors.New(GetCmErr + err.Error()))
					}
					for k, ve := range cm2.Data {
						cm.CmData[k] = ve
					}
					rcm.ResourceConfigmap = append(rcm.ResourceConfigmap, cm)
					cm = p.Configmap{}
				}
			}
		}
		acm.AllConfigmap = append(acm.AllConfigmap, rcm)
		rcm = p.ResourceConfigmap{}
	}
	tools.Writer.WriteIntoFile(config.ConfigmapFile, &acm)
	filePath := config.DataDirectory + config.ConfigmapFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
