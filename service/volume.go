package service

import (
	b "k8s-baseline-scanner-v2/base"
)

func PullAllVolume() {
	var (
		rceSce  []b.DeployStsDsBase
		allvol  b.AllVolume
		rvol    b.ResourceVolume
		msgName = "volume"
	)
	allvol = make(map[string]b.ResourceVolume)
	rceSce = b.GetResources()
	for _, rce := range rceSce {
		for _, vol := range rce.Volumes {
			rvol = append(rvol, vol.Name)
		}
		allvol[rce.ResourceInfo] = rvol
		rvol = b.ResourceVolume{}
	}
	b.WriteDataIntoFile(msgName, b.VolumeFile, &allvol)
}
