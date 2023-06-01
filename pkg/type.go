package pkg

import (
	v1 "k8s.io/api/core/v1"
)

// 所有configmap
type AllConfigmap struct {
	AllConfigmap []ResourceConfigmap `yaml:"allConfigmap"`
}

// 单个应用（deploy、sts、ds）对应的configmap
type ResourceConfigmap struct {
	ResourceInfo      string      `yaml:"resourceInfo"`
	ResourceConfigmap []Configmap `yaml:"resourceConfigmap,omitempty"`
}

// 单个configmap
type Configmap struct {
	CmName string            `yaml:"cmName,omitempty"`
	CmData map[string]string `yaml:"cmData,omitempty"`
}

// 存储卷挂载
type AllVolumeMount struct {
	AllVolumeMount []ResourceVolumeMount `yaml:"allVolumeMount"`
}

type ResourceVolumeMount struct {
	ResourceInfo        string                 `yaml:"resourceInfo"`
	ResourceVolumeMount []ContainerVolumeMount `yaml:"resourceVolumeMount,omitempty"`
}

// 单个容器存储卷挂载
type ContainerVolumeMount struct {
	ContainerName string            `yaml:"containerName,omitempty"`
	VolumeMount   map[string]string `yaml:"volumeMount,omitempty"`
}

// 所有应用的容忍
type AllToleration struct {
	AllToleration []ResourceToleration `yaml:"allToleration"`
}

// 单个应用（deploy、sts、ds）对应的容忍
type ResourceToleration struct {
	ResourceInfo       string   `yaml:"resourceInfo"`
	ResourceToleration []string `yaml:"resourceToleration,omitempty"`
}

// 单个应用（deploy、sts、ds）对应的节点选择器
type ResourceNodeSelector struct {
	ResourceInfo     string            `yaml:"resourceInfo"`
	ResourceSelector map[string]string `yaml:"resourceSelector,omitempty"`
}

// 所有的节点选择器
type AllNodeSelector struct {
	AllNodeSelector []ResourceNodeSelector `yaml:"allNodeSelector"`
}

// 所有pvc
type Allpvc struct {
	AllPvc []Pvc `yaml:"allPvc"`
}

// 单个应用（deploy、sts、ds）对应的pvc
type Pvc struct {
	PvcInfo []string `yaml:"pvcInfo"`
}

// 单个应用（deploy、sts、ds）对应的亲和,反亲和
type ResourceAffinity struct {
	ResourceInfo     string       `yaml:"resourceInfo"`
	ResourceAffinity *v1.Affinity `yaml:"resourceAffinity,omitempty"`
}

// 所有亲和、反亲和
type AllAffinity struct {
	AllAffinity []ResourceAffinity `yaml:"allAffinity"`
}

// 应用的资源
type AllResource struct {
	AllResource []DeployStsDsResource `yaml:"allResource"`
}
type DeployStsDsResource struct {
	DeployStsDsResource []map[string]string `yaml:"deployStsDsResource"`
}

// 镜像
type AllImage struct {
	AllImage []ResourceImage `yaml:"allImage"`
}
type ResourceImage struct {
	ResourceInfo  string              `yaml:"resourceInfo"`
	ResourceImage []map[string]string `yaml:"resourceImage"`
}

// 所有探针
type AllProbe struct {
	AllProbe []ResourceProbe `yaml:"allprobe"`
}

// 单个应用包含的探针
type ResourceProbe struct {
	ResourceInfo  string           `yaml:"resourceInfo,omitempty"`
	ResourceProbe []ContainerProbe `yaml:"resourceProbe,omitempty"`
}

// 单个容器包含的探针
type ContainerProbe struct {
	ContainerName  string    `yaml:"containerName,omitempty"`
	LivenessProbe  ProbeRule `yaml:"livenessProbe,omitempty"`
	ReadinessProbe ProbeRule `yaml:"redinessProbe,omitempty"`
	StartUpProbe   ProbeRule `yaml:"startUpProbe,omitempty"`
}

// 单个探针的内容
type ProbeRule struct {
	HTTPGet       string `yaml:"httpGet,omitempty"`
	TcpSocket     string `yaml:"tcpSocket,omitempty"`
	Command       string `yaml:"command,omitempty"`
	ProbeStretegy string `yaml:"probeStretegy,omitempty"`
}

// 所有的容器信息
type AllContainers struct {
	AllContainers []ResourceContainer `yaml:"allContainers,omitempty"`
}

// 单个应用（deploy、sts、ds）对应的容器信息
type ResourceContainer struct {
	ResourceInfo string         `yaml:"info"`
	Containers   []v1.Container `yaml:"containers,omitempty"`
}

// deployment、statefulset、daemonset 三种的基础结构
type DeployStsDsBase struct {
	ResourceInfo       string            `yaml:"resourceInfo"`
	Replicas           int32             `yaml:"replicas,omitempty"`
	Volumes            []v1.Volume       `yaml:"volumes,omitempty"`
	RestartPolicy      string            `yaml:"restartPolicy,omitempty"`
	DnsPolicy          string            `yaml:"dnsPolicy,omitempty"`
	NodeSelector       map[string]string `yaml:"nodeSelector,omitempty"`
	ServiceAccountName string            `yaml:"serviceAccountName,omitempty"`
	Affinity           *v1.Affinity      `yaml:"affinity,omitempty"`
	Tolerations        []v1.Toleration   `yaml:"tolerations,omitempty"`
}

// 所有volume
type AllVolume struct {
	AllVolume []ResourceVolume `yaml:"allVolume"`
}

// 单个应用（deploy、sts、ds）对应的volume
type ResourceVolume struct {
	ResourceInfo    string   `yaml:"resourceInfo"`
	ResourceVolumes []string `yaml:"resourceVolumes,omitempty"`
}

// 副本数
type AllReplicas struct {
	AllReplicas []ResourceReplicas `yaml:"allReplicas"`
}

type ResourceReplicas struct {
	ResourceInfo     string `yaml:"resourceInfo"`
	ResourceReplicas string `yaml:"resourceReplicas"`
}

// 所有secret
type AllSecret struct {
	AllSecretInfo []ResourceSecret `yaml:"allSecretInfo"`
}

// 单个应用（deploy、sts、ds）对应的secret
type ResourceSecret struct {
	ResourceInfo string   `yaml:"resourceInfo"`
	SecretNames  []string `yaml:"secretNames,omitempty"`
}

// 所有service
type AllService struct {
	AllService []Service `yaml:"allService"`
}

// 单个service
type Service struct {
	ServiceInfo string            `yaml:"serviceInfo"`
	Selector    map[string]string `yaml:"selector,omitempty"`
	Potrs       []string          `yaml:"potrs,omitempty"`
}

// 所有env
type AllEnv struct {
	AllEnv []ResourceEnv `yaml:"allEnv"`
}

// 单个应用（deploy、sts、ds）对应的环境变量
type ResourceEnv struct {
	ResourceInfo string         `yaml:"resourceInfo"`
	ResourceEnv  []ContainerEnv `yaml:"resourceEnv,omitempty"`
}

// 单个容器的环境变量
type ContainerEnv struct {
	ContainerName string             `yaml:"containerName"`
	Env           []v1.EnvVar        `yaml:"env,omitempty"`
	EnvFrom       []v1.EnvFromSource `yaml:"envFrom,omitempty"`
}

type AllContainerResource struct {
	AllContainerResource []RContainerResource `yaml:"allContainerResource"`
}

// 单个应用（deploy、sts、ds）对应的容器资源
type RContainerResource struct {
	ResourceInfo       string              `yaml:"resourceInfo"`
	RContainerResource []ContainerResource `yaml:"ContainerResource"`
}

// 单个容器资源规格
type ContainerResource struct {
	ContainerName string   `yaml:"containerName"`
	Request       string   `yaml:"request,omitempty"`
	Limit         string   `yaml:"limit,omitempty"`
	Claims        []string `yaml:"claims,omitempty"`
}
