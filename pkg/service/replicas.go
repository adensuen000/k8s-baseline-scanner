package service

import (
	"k8s-baseline-scanner/config"
	p "k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
	"strconv"
)

func (*resource) GetAllReplicas() {
	var (
		rceSce []p.DeployStsDsBase
		rr     p.ResourceReplicas
		ar     p.AllReplicas
		msg    = "replicas"
	)
	rceSce = p.GetResources()
	for _, rce := range rceSce {
		rr.ResourceInfo = rce.ResourceInfo
		rr.ResourceReplicas = strconv.Itoa(int(rce.Replicas))
		ar.AllReplicas = append(ar.AllReplicas, rr)
	}
	tools.Writer.WriteIntoFile(config.ReplicasFile, &ar)
	filePath := config.DataDirectory + config.ReplicasFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
