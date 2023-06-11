package service

import (
	"k8s-baseline-scanner-v2/base"
)

const (
	curImageFile = base.DataDirectory + base.ImageFile
	stdImageFile = base.StdBaselineDir + base.ImageFile

	curRepFile = base.DataDirectory + base.ReplicasFile
	stdRepFile = base.StdBaselineDir + base.ReplicasFile

	curPvcFile = base.DataDirectory + base.PvcFile
	stdPvcFile = base.StdBaselineDir + base.PvcFile

	curSvcFile = base.DataDirectory + base.ServiceFile
	stdSvcFile = base.StdBaselineDir + base.ServiceFile

	curSecretFile = base.DataDirectory + base.SecretFile
	stdSecretFile = base.StdBaselineDir + base.SecretFile

	curVolumeFile = base.DataDirectory + base.VolumeFile
	stdVolumeFile = base.StdBaselineDir + base.VolumeFile

	curCmFile = base.DataDirectory + base.ConfigmapFile
	stdCmFile = base.StdBaselineDir + base.ConfigmapFile

	curAffinityFile = base.DataDirectory + base.AffinityFile
	stdAffinityFile = base.StdBaselineDir + base.AffinityFile

	curContainerResourceFile = base.DataDirectory + base.ContainerResourceFile
	stdContainerResourceFile = base.StdBaselineDir + base.ContainerResourceFile

	curNodeSelectorFile = base.DataDirectory + base.NodeSelectorFile
	stdNodeSelectorFile = base.StdBaselineDir + base.NodeSelectorFile

	curTolerationFile = base.DataDirectory + base.TolerationFile
	stdTolerationFile = base.StdBaselineDir + base.TolerationFile

	curEnvFile = base.DataDirectory + base.EnvFile
	stdEnvFile = base.StdBaselineDir + base.EnvFile

	curProbeFile = base.DataDirectory + base.ProbeFile
	stdProbeFile = base.StdBaselineDir + base.ProbeFile

	curVolumeMountFile = base.DataDirectory + base.VolumeMountFile
	stdVolumeMountFile = base.StdBaselineDir + base.VolumeMountFile
)

// 输出image比对结果
func GetImageCompareRes() {
	msg := "image"
	ResList := compareResource(curImageFile, stdImageFile)
	getCompareRes(msg, ResList)
}

// 输出replicas比对结果
func GetReplicasCompareRes() {
	msg := "replicas"
	ResList := compareResource(curRepFile, stdRepFile)
	getCompareRes(msg, ResList)
}

// 输出pvc比对结果
func GetPvcCompareRes() {
	msg := "pvc"
	ResList := compareResource(curPvcFile, stdPvcFile)
	getCompareRes(msg, ResList)
}

func GetServiceCompareRes() {
	msg := "service"
	ResList := compareResource(curSvcFile, stdSvcFile)
	getCompareRes(msg, ResList)
}

func GetSecretCompareRes() {
	msg := "secret"
	ResList := compareResource(curSecretFile, stdSecretFile)
	getCompareRes(msg, ResList)
}

func GetvolumeCompareRes() {
	msg := "volume"
	ResList := compareResource(curVolumeFile, stdVolumeFile)
	getCompareRes(msg, ResList)
}

func GetConfigmapCompareRes() {
	msg := "configmap"
	ResList := compareResource(curCmFile, stdCmFile)
	getCompareRes(msg, ResList)
}

func GetAffinityCompareRes() {
	msg := "affinity"
	ResList := compareResource(curAffinityFile, stdAffinityFile)
	getCompareRes(msg, ResList)
}

func GetContainerResourceCompareRes() {
	msg := "affinity"
	ResList := compareResource(curContainerResourceFile, stdContainerResourceFile)
	getCompareRes(msg, ResList)
}

func GetNodeSelectorCompareRes() {
	msg := "nodeSelector"
	ResList := compareResource(curNodeSelectorFile, stdNodeSelectorFile)
	getCompareRes(msg, ResList)

}

func GetTolerationCompareRes() {
	msg := "nodeSelector"
	ResList := compareResource(curTolerationFile, stdTolerationFile)
	getCompareRes(msg, ResList)

}

func GetEnvCompareRes() {
	msg := "env"
	ResList := compareResource(curEnvFile, stdEnvFile)
	getCompareRes(msg, ResList)

}

func GetProbeCompareRes() {
	msg := "probe"
	ResList := compareResource(curProbeFile, stdProbeFile)
	getCompareRes(msg, ResList)

}

func GetVolumeMountCompareRes() {
	msg := "VolumeMount"
	ResList := compareResource(curVolumeMountFile, stdVolumeMountFile)
	getCompareRes(msg, ResList)

}

func CompareAll() {
	GetServiceCompareRes()
	GetSecretCompareRes()
	GetPvcCompareRes()
	GetvolumeCompareRes()
	GetConfigmapCompareRes()
	GetReplicasCompareRes()
	GetAffinityCompareRes()
	GetContainerResourceCompareRes()
	GetNodeSelectorCompareRes()
	GetTolerationCompareRes()
	GetEnvCompareRes()
	GetImageCompareRes()
	GetProbeCompareRes()
	GetVolumeMountCompareRes()
}
