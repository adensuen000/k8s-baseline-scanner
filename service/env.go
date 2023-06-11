package service

import (
	b "k8s-baseline-scanner-v2/base"
)

func PullAllEnv() {
	var (
		cEnv    b.ContainerEnv
		rEnv    b.ResourceEnv
		aEnv    b.AllEnv
		msgName = "env"
		file    = b.EnvFile
	)
	aEnv = make(map[string]b.ResourceEnv)
	for _, rce := range b.K8sRce.GetContainers() {
		aEnvKey := rce.ResourceInfo
		for _, ctr := range rce.Containers {
			cEnv.ContainerName = ctr.Name
			cEnv.Env = ctr.Env
			cEnv.EnvFrom = ctr.EnvFrom
			rEnv = append(rEnv, cEnv)
			cEnv = b.ContainerEnv{}
		}
		aEnv[aEnvKey] = rEnv
		rEnv = b.ResourceEnv{}
	}
	b.WriteDataIntoFile(msgName, file, &aEnv)
}
