package service

import (
	"fmt"
	b "k8s-baseline-scanner-v2/base"
	"strconv"
)

func PullAllToleration() {
	var (
		tol     b.ResourceToleration
		atol    b.AllToleration
		rceSce  []b.DeployStsDsBase
		tolInfo string
		msgName = "toleration"
		file    = b.TolerationFile
	)
	atol = make(map[string]b.ResourceToleration)
	rceSce = b.GetResources()
	for _, rce := range rceSce {
		atolKey := rce.ResourceInfo
		for _, tole := range rce.Tolerations {
			if tole.TolerationSeconds == nil {
				tolInfo = tole.Key + "," + tole.Value + "," + fmt.Sprintf(string(tole.Effect)) + ","
			} else {
				tolInfo = tole.Key + "," + tole.Value + "," + string(tole.Effect) + "," + strconv.FormatInt(*tole.TolerationSeconds, 10)
			}
			tol = append(tol, tolInfo)
		}
		atol[atolKey] = tol
		tol = b.ResourceToleration{}
	}
	b.WriteDataIntoFile(msgName, file, &atol)
}
