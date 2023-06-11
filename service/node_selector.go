package service

import (
	b "k8s-baseline-scanner-v2/base"
)

func PullAllNodeSelector() {
	var (
		rsel    b.ResourceNodeSelector
		asel    b.AllNodeSelector
		msgName = "nodeSelector"
		file    = b.NodeSelectorFile
	)
	asel = make(map[string]b.ResourceNodeSelector)
	for _, rce := range b.GetResources() {
		rsel = make(map[string]string)
		aselKey := rce.ResourceInfo
		for k, v := range rce.NodeSelector {
			rsel[k] = v
		}
		asel[aselKey] = rsel
		rsel = b.ResourceNodeSelector{}
	}
	b.WriteDataIntoFile(msgName, file, &asel)
}
