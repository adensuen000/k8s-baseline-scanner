package service

import (
	"k8s-baseline-scanner/config"
	"k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
)

func (*resource) GetAllContainerResource() {
	var (
		crs    pkg.ContainerResource
		rcrs   pkg.RContainerResource
		allcrs pkg.AllContainerResource
		msg    = "container-resource"
	)
	for _, actrs := range pkg.K8sRce.GetContainers().AllContainers {
		rcrs.ResourceInfo = actrs.ResourceInfo
		for _, ctr := range actrs.Containers {
			crs.ContainerName = ctr.Name
			crs.Request = ctr.Resources.Requests.Cpu().String() + "," + ctr.Resources.Requests.Memory().String()
			crs.Limit = ctr.Resources.Limits.Cpu().String() + "," + ctr.Resources.Limits.Memory().String()
			for _, claim := range ctr.Resources.Claims {
				crs.Claims = append(crs.Claims, claim.Name)
			}
			rcrs.RContainerResource = append(rcrs.RContainerResource, crs)
			crs = pkg.ContainerResource{}
		}
		allcrs.AllContainerResource = append(allcrs.AllContainerResource, rcrs)
		rcrs = pkg.RContainerResource{}
	}
	tools.Writer.WriteIntoFile(config.ContainerResourceFile, &allcrs)
	filePath := config.DataDirectory + config.ContainerResourceFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
