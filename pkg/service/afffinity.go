package service

import (
	"k8s-baseline-scanner/config"
	p "k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
)

func (*resource) GetAllAffinity() {
	var (
		rceAff p.ResourceAffinity
		allAff p.AllAffinity
		msg    = "affinity"
	)
	for _, rce := range p.GetResources() {
		rceAff.ResourceInfo = rce.ResourceInfo
		rceAff.ResourceAffinity = rce.Affinity
		allAff.AllAffinity = append(allAff.AllAffinity, rceAff)
		rceAff = p.ResourceAffinity{}
	}
	tools.Writer.WriteIntoFile(config.AffinityFile, &allAff)
	filePath := config.DataDirectory + config.AffinityFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
