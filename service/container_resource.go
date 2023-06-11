package service

import (
	b "k8s-baseline-scanner-v2/base"
)

func PullAllContainerResources() {
	var (
		crs     b.ContainerResource
		rcrs    b.RContainerResource
		allcrs  b.AllContainerResources
		msgName = "containerResource"
		file    = b.ContainerResourceFile
	)
	allcrs = make(map[string]b.RContainerResource)
	for _, actrs := range b.K8sRce.GetContainers() {
		allcrsKey := actrs.ResourceInfo
		for _, ctr := range actrs.Containers {
			crs.ContainerName = ctr.Name
			crs.ReqLim = ctr.Resources.Requests.Cpu().String() + "," +
				ctr.Resources.Requests.Memory().String() + "," +
				ctr.Resources.Limits.Cpu().String() + "," +
				ctr.Resources.Limits.Memory().String()
			for _, claim := range ctr.Resources.Claims {
				crs.Claims = append(crs.Claims, claim.Name)
			}
			rcrs.RContainerResource = append(rcrs.RContainerResource, crs)
			crs = b.ContainerResource{}
		}
		allcrs[allcrsKey] = rcrs
		rcrs = b.RContainerResource{}
	}
	b.WriteDataIntoFile(msgName, file, &allcrs)
}
