package service

import (
	"k8s-baseline-scanner/config"
	"k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
)

func (*container) GetAllImage() {
	var (
		rceImg pkg.ResourceImage
		allImg pkg.AllImage
		image  map[string]string
		msg    = "image"
	)
	for _, actrs := range pkg.K8sRce.GetContainers().AllContainers {
		rceImg.ResourceInfo = actrs.ResourceInfo
		for _, ctr := range actrs.Containers {
			image = make(map[string]string, 10)
			image[ctr.Name] = ctr.Image
			rceImg.ResourceImage = append(rceImg.ResourceImage, image)
			image = nil
		}
		allImg.AllImage = append(allImg.AllImage, rceImg)
		rceImg = pkg.ResourceImage{}
	}
	tools.Writer.WriteIntoFile(config.ImageFile, &allImg)
	filePath := config.DataDirectory + config.ImageFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
