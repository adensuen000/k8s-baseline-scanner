package service

import (
	b "k8s-baseline-scanner-v2/base"
	"strconv"
)

func PullAllReplica() {
	var (
		rceSce  []b.DeployStsDsBase
		ar      b.AllReplica
		msgName = "replicas"
	)
	rceSce = b.GetResources()
	for _, rce := range rceSce {
		ar[rce.ResourceInfo] = strconv.Itoa(int(rce.Replicas))
	}
	b.WriteDataIntoFile(msgName, b.ReplicasFile, &ar)
}
