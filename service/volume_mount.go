package service

import (
	b "k8s-baseline-scanner-v2/base"
)

func PullAllVolumeMount() {
	var (
		vm      b.VolumeMount
		vmlist  []b.VolumeMount
		cvm     b.ContainerVolumeMount
		rvm     b.ResourceVolumeMount
		avm     b.AllVolumeMount
		msgName = "volume-mount"
		file    = b.VolumeMountFile
	)
	avm = make(map[string]b.ResourceVolumeMount)
	for _, actrs := range b.K8sRce.GetContainers() {
		for _, ctr := range actrs.Containers {
			cvm = make(map[string][]b.VolumeMount)
			cvmKey := ctr.Name
			for _, volm := range ctr.VolumeMounts {
				vm.Name, vm.MountPath, vm.ReadOnly, vm.SubPath, vm.SubPathExpr, vm.MountPropagation = volm.Name, volm.MountPath, volm.ReadOnly, volm.SubPath, volm.SubPathExpr, volm.MountPropagation
				vmlist = append(vmlist, vm)
			}
			cvm[cvmKey] = vmlist
			vmlist = nil
			rvm = append(rvm, cvm)
			cvm = nil
		}
		avmKey := actrs.ResourceInfo
		avm[avmKey] = rvm
		rvm = b.ResourceVolumeMount{}
	}
	b.WriteDataIntoFile(msgName, file, &avm)
}
