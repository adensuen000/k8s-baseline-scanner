package service

import (
	b "k8s-baseline-scanner-v2/base"
)

func PullAllImage() {
	var (
		allImg  b.AllImage
		ri      b.ResourceImage
		msgName = "image"
		file    = b.ImageFile
	)
	allImg = make(map[string]b.ResourceImage)
	for _, actrs := range b.K8sRce.GetContainers() {
		ri = make(map[string]string)
		key, containerList := actrs.ResourceInfo, actrs.Containers
		for _, ctr := range containerList {
			ri[ctr.Name] = ctr.Image
		}
		allImg[key] = ri
		ri = b.ResourceImage{}
	}
	b.WriteDataIntoFile(msgName, file, &allImg)
}
