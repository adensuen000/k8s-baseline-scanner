package service

import (
	"k8s-baseline-scanner/config"
	p "k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
)

func (*resource) GetAllVolume() {
	var (
		rceSce []p.DeployStsDsBase
		rvol   p.ResourceVolume
		allvol p.AllVolume
		msg    = "volume"
	)
	rceSce = p.GetResources()
	for _, rce := range rceSce {
		for _, vol := range rce.Volumes {
			rvol.ResourceVolumes = append(rvol.ResourceVolumes, vol.Name)
		}
		rvol.ResourceInfo = rce.ResourceInfo
		allvol.AllVolume = append(allvol.AllVolume, rvol)
		rvol = p.ResourceVolume{}
	}
	tools.Writer.WriteIntoFile(config.VolumeFile, &allvol)
	filePath := config.DataDirectory + config.VolumeFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
