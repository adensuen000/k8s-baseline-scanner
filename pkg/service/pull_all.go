package service

func StartAllAction() {
	Container.GetAllImage()
	Container.GetAllVolumeMount()
	Container.GetAllProbe()
	Resource.GetAllVolume()
	Resource.GetAllConfigmap()
	Resource.GetAllToleration()
	Resource.GetAllNodeSelector()
	Resource.GetAllAffinity()
	Resource.GetAllPvc()
	Resource.GetAllReplicas()
	Resource.GetAllService()
	Resource.GetAllSecret()
	Container.GetAllEnv()
	Resource.GetAllContainerResource()
}
