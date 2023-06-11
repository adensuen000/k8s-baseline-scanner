package service

import (
	b "k8s-baseline-scanner-v2/base"
)

func PullAllSecret() {
	var (
		asec    b.AllSecret
		rsec    b.ResourceSecret
		msgName = "secret"
		file    = b.SecretFile
	)
	asec = make(map[string]b.ResourceSecret)
	for _, rce := range b.GetResources() {
		asecKey := rce.ResourceInfo
		for _, sec := range rce.Volumes {
			if sec.Secret != nil {
				rsec = append(rsec, sec.Secret.SecretName)
			}
		}
		asec[asecKey] = rsec
		rsec = b.ResourceSecret{}
	}
	b.WriteDataIntoFile(msgName, file, &asec)
}
