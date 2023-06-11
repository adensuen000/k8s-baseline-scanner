package service

import (
	b "k8s-baseline-scanner-v2/base"
	v1 "k8s.io/api/core/v1"
)

func PullAllAffinity() {
	var (
		allAff  b.AllAffinity
		msgName = "affinity"
		file    = b.AffinityFile
	)
	allAff = make(map[string]*v1.Affinity)
	for _, rce := range b.GetResources() {
		allAffKey := rce.ResourceInfo
		allAff[allAffKey] = rce.Affinity
	}
	b.WriteDataIntoFile(msgName, file, &allAff)
}
