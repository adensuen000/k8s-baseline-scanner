package service

func PullAllResource() {
	PullAllService()
	PullAllSecret()
	PullAllPvc()
	PullAllVolume()
	PullAllConfigmap()
	PullAllReplica()
	PullAllAffinity()
	PullAllContainerResources()
	PullAllNodeSelector()
	PullAllToleration()
	PullAllEnv()
	PullAllImage()
	PullAllProbe()
	PullAllVolumeMount()
}
