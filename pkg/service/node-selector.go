package service

import (
	"k8s-baseline-scanner/config"
	p "k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
)

func (*resource) GetAllNodeSelector() {
	var (
		rsel p.ResourceNodeSelector
		asel p.AllNodeSelector
		msg  = "node-selector"
	)
	for _, rce := range p.GetResources() {
		rsel.ResourceSelector = make(map[string]string)
		rsel.ResourceInfo = rce.ResourceInfo
		for k, v := range rce.NodeSelector {
			rsel.ResourceSelector[k] = v
		}
		asel.AllNodeSelector = append(asel.AllNodeSelector, rsel)
		rsel = p.ResourceNodeSelector{}
	}
	tools.Writer.WriteIntoFile(config.NodeSelectorFile, &asel)
	filePath := config.DataDirectory + config.NodeSelectorFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
