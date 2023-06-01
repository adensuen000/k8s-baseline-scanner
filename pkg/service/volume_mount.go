package service

import (
	"k8s-baseline-scanner/config"
	"k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
)

func (*container) GetAllVolumeMount() {
	var (
		cvm pkg.ContainerVolumeMount
		rvm pkg.ResourceVolumeMount
		avm pkg.AllVolumeMount
		msg = "volume-mount"
	)
	for _, actrs := range pkg.K8sRce.GetContainers().AllContainers {
		for _, ctr := range actrs.Containers {
			for _, volm := range ctr.VolumeMounts {
				cvm.VolumeMount = make(map[string]string)
				cvm.VolumeMount[volm.Name] = volm.MountPath
			}
			cvm.ContainerName = ctr.Name
			rvm.ResourceVolumeMount = append(rvm.ResourceVolumeMount, cvm)
			cvm = pkg.ContainerVolumeMount{}
		}
		rvm.ResourceInfo = actrs.ResourceInfo
		avm.AllVolumeMount = append(avm.AllVolumeMount, rvm)
		rvm = pkg.ResourceVolumeMount{}
	}
	tools.Writer.WriteIntoFile(config.VolumeMountFile, &avm)
	filePath := config.DataDirectory + config.VolumeMountFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
