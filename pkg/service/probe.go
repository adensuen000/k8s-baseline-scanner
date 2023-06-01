package service

import (
	"fmt"
	"k8s-baseline-scanner/config"
	"k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/tools"
	v1 "k8s.io/api/core/v1"
)

func getProbeRule(probe *v1.Probe) *pkg.ProbeRule {
	pr := &pkg.ProbeRule{}
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

func (*container) GetAllProbe() {
	var (
		ctrProbe pkg.ContainerProbe
		rceProbe pkg.ResourceProbe
		allProbe pkg.AllProbe
		msg      = "probe"
	)
	for _, actrs := range pkg.K8sRce.GetContainers().AllContainers {
		rceProbe.ResourceInfo = actrs.ResourceInfo
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
			rceProbe.ResourceProbe = append(rceProbe.ResourceProbe, ctrProbe)
			ctrProbe = pkg.ContainerProbe{}
		}
		allProbe.AllProbe = append(allProbe.AllProbe, rceProbe)
		rceProbe = pkg.ResourceProbe{}
	}
	tools.Writer.WriteIntoFile(config.ProbeFile, &allProbe)
	filePath := config.DataDirectory + config.ProbeFile
	tools.CleanNullInFile(filePath)
	tools.GetSuccessMsg(msg)
}
