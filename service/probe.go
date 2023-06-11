package service

import (
	"fmt"
	b "k8s-baseline-scanner-v2/base"
	v1 "k8s.io/api/core/v1"
)

func PullAllProbe() {
	var (
		ctrProbe b.ContainerProbe
		rceProbe b.ResourceProbe
		allProbe b.AllProbe
		msgName  = "probe"
		file     = b.ProbeFile
	)
	allProbe = make(map[string]b.ResourceProbe)
	for _, actrs := range b.K8sRce.GetContainers() {
		allProbeKey := actrs.ResourceInfo
		for _, ctr := range actrs.Containers {
			ctrProbe.ContainerName = ctr.Name
			if ctr.LivenessProbe != nil {
				ctrProbe.LivenessProbe = *getProbeRule(ctr.LivenessProbe)
			}
			if ctr.ReadinessProbe != nil {
				ctrProbe.ReadinessProbe = *getProbeRule(ctr.ReadinessProbe)
			}
			if ctr.StartupProbe != nil {
				ctrProbe.StartUpProbe = *getProbeRule(ctr.StartupProbe)
			}
			rceProbe = append(rceProbe, ctrProbe)
			ctrProbe = b.ContainerProbe{}
		}
		allProbe[allProbeKey] = rceProbe
		rceProbe = b.ResourceProbe{}
	}
	b.WriteDataIntoFile(msgName, file, &allProbe)
}

func getProbeRule(probe *v1.Probe) *b.ProbeRule {
	pr := &b.ProbeRule{}
	probeStretegy := fmt.Sprintf("InitialDelaySeconds:%d,PeriodSeconds:%d,TimeoutSeconds:%d,FailureThreshold:%d,SuccessThreshold:%d",
		probe.InitialDelaySeconds,
		probe.PeriodSeconds,
		probe.TimeoutSeconds,
		probe.FailureThreshold,
		probe.SuccessThreshold)
	pr.ProbeStretegy = probeStretegy

	if probe.HTTPGet != nil {
		pr.HTTPGet = fmt.Sprintf("Scheme:%s,Port:%s,Path:%s", probe.HTTPGet.Scheme, probe.HTTPGet.Port.String(), probe.HTTPGet.Path)
	}
	if probe.TCPSocket != nil {
		pr.TcpSocket = fmt.Sprintf("HOST:%s,Port:%s", probe.TCPSocket.Host, probe.TCPSocket.Port.String())

	}
	if probe.Exec != nil {
		pr.Command = fmt.Sprintf("%s", probe.Exec.Command)
	}
	return pr
}
