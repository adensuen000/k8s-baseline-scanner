package service

import (
	"k8s-baseline-scanner/config"
	p "k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
)

func (*resource) GetAllSecret() {
	var (
		asec p.AllSecret
		rsec p.ResourceSecret
		msg  = "secret"
	)
	for _, rce := range p.GetResources() {
		rsec.ResourceInfo = rce.ResourceInfo
		for _, sec := range rce.Volumes {
			if sec.Secret != nil {
				rsec.SecretNames = append(rsec.SecretNames, sec.Secret.SecretName)
			}
		}
		asec.AllSecretInfo = append(asec.AllSecretInfo, rsec)
		rsec = p.ResourceSecret{}
	}
	tools.Writer.WriteIntoFile(config.SecretFile, &asec)
	filePath := config.DataDirectory + config.SecretFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
