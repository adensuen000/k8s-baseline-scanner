package service

import (
	"k8s-baseline-scanner/config"
	p "k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
)

func (*container) GetAllEnv() {
	var (
		cEnv p.ContainerEnv
		rEnv p.ResourceEnv
		aEnv p.AllEnv
		msg  = "env"
	)

	for _, rce := range p.K8sRce.GetContainers().AllContainers {
		rEnv.ResourceInfo = rce.ResourceInfo
		for _, ctr := range rce.Containers {
			cEnv.ContainerName = ctr.Name
			cEnv.Env = ctr.Env
			cEnv.EnvFrom = ctr.EnvFrom
			rEnv.ResourceEnv = append(rEnv.ResourceEnv, cEnv)
			cEnv = p.ContainerEnv{}
		}
		aEnv.AllEnv = append(aEnv.AllEnv, rEnv)
		rEnv = p.ResourceEnv{}
	}
	tools.Writer.WriteIntoFile(config.EnvFile, &aEnv)
	filePath := config.DataDirectory + config.EnvFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
