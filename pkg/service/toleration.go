package service

import (
	"fmt"
	"k8s-baseline-scanner/config"
	p "k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
	"strconv"
)

func (*resource) GetAllToleration() {
	var (
		tol     p.ResourceToleration
		atol    p.AllToleration
		rceSce  []p.DeployStsDsBase
		tolInfo string
		msg     = "toleration"
	)
	rceSce = p.GetResources()
	for _, rce := range rceSce {
		tol.ResourceInfo = rce.ResourceInfo
		for _, tole := range rce.Tolerations {
			if tole.TolerationSeconds == nil {
				tolInfo = tole.Key + "," + tole.Value + "," + fmt.Sprintf(string(tole.Effect)) + ","
			} else {
				tolInfo = tole.Key + "," + tole.Value + "," + string(tole.Effect) + "," + strconv.FormatInt(*tole.TolerationSeconds, 10)
			}
			tol.ResourceToleration = append(tol.ResourceToleration, tolInfo)
		}
		atol.AllToleration = append(atol.AllToleration, tol)
		tol = p.ResourceToleration{}
	}
	tools.Writer.WriteIntoFile(config.TolerationFile, &atol)
	filePath := config.DataDirectory + config.TolerationFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
